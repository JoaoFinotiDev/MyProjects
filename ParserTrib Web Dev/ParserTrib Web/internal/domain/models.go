package domain

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// TipoAcao define os tipos de ações disponíveis no menu
type TipoAcao int

const (
	AcaoArquivo TipoAcao = iota + 1
	AcaoRefresh
	AcaoSair
)

// ResultadoMenu representa o resultado do menu atribuindo o resultado ao tipo de ação correspondente
type ResultadoMenu struct {
	Acao    TipoAcao
	Arquivo ArquivoExcel
}

// ArquivoExcel representa a estrutura de informações do arquivo Excel
type ArquivoExcel struct {
	Nome            string
	Caminho         string
	DataModificacao time.Time
}

// Planilha representa metadados da planilha do Excel
type Planilha struct {
	Caminho     string
	NomeSheet   string
	Cabecalhos  []string
	TotalLinhas int
}

// ErroValidacao representa um erro especifico
type ErroValidacao struct {
	Linha      int    `json:"linha"`
	Coluna     string `json:"coluna"`
	NomeColuna string `json:"nomeColuna"`
	Tipo       string `json:"tipo"`
	Mensagem   string `json:"mensagem"`
}

// String formatada de mensagem de erro para exibição
func (e ErroValidacao) String() string {
	return fmt.Sprintf("[ERRO] Linha %d, Coluna %s (%s): %s",
		e.Linha,
		e.Coluna,
		e.NomeColuna,
		e.Mensagem)
}

// ResultadoValidacaoCompleto agrupa os 4 tipos de erros
type ResultadoValidacaoCompleto struct {
	NomeArquivo    string          `json:"nomeArquivo"`
	ErrosVazias    []ErroValidacao `json:"errosVazias"`
	ErrosNCM       []ErroValidacao `json:"errosNCM"`
	ErrosCSTOrigem []ErroValidacao `json:"errosCSTOrigem"`
	ErrosCSOSN     []ErroValidacao `json:"errosCSOSN"`
	ErrosTipoItem  []ErroValidacao `json:"errosTipoItem"`
	TempoExecucao  time.Duration   `json:"-"`
}

// RespostaValidacaoAPI é a estrutura serializada para a API
type RespostaValidacaoAPI struct {
	NomeArquivo    string          `json:"nomeArquivo"`
	TempoExecucao  string          `json:"processingTime"`
	TotalErros     int             `json:"totalErros"`
	ErrosVazias    int             `json:"errosVazias"`
	ErrosNCM       int             `json:"errosNCM"`
	ErrosCSTOrigem int             `json:"errosCSTOrigem"`
	ErrosCSOSN     int             `json:"errosCSOSN"`
	ErrosTipoItem  int             `json:"errosTipoItem"`
	Detalhes       []ErroValidacao `json:"detalhes"`
}

// MarshalJSON custom para RespostaValidacaoAPI não é necessário — campos são primitivos

// ToRespostaAPI converte ResultadoValidacaoCompleto para RespostaValidacaoAPI
func (r ResultadoValidacaoCompleto) ToRespostaAPI() RespostaValidacaoAPI {
	detalhes := make([]ErroValidacao, 0)

	for _, e := range r.ErrosVazias {
		e.Tipo = "VAZIA"
		detalhes = append(detalhes, e)
	}
	for _, e := range r.ErrosNCM {
		e.Tipo = "NCM"
		detalhes = append(detalhes, e)
	}
	for _, e := range r.ErrosCSTOrigem {
		e.Tipo = "CST_ORIGEM"
		detalhes = append(detalhes, e)
	}
	for _, e := range r.ErrosCSOSN {
		e.Tipo = "CSOSN"
		detalhes = append(detalhes, e)
	}
	for _, e := range r.ErrosTipoItem {
		e.Tipo = "TIPO_ITEM"
		detalhes = append(detalhes, e)
	}

	// Ordenar por coluna (alfabética) e depois por linha (numérica)
	sort.Slice(detalhes, func(i, j int) bool {
		if detalhes[i].Coluna != detalhes[j].Coluna {
			return detalhes[i].Coluna < detalhes[j].Coluna
		}
		return detalhes[i].Linha < detalhes[j].Linha
	})

	return RespostaValidacaoAPI{
		NomeArquivo:    r.NomeArquivo,
		TempoExecucao:  r.TempoExecucao.String(),
		TotalErros:     r.TotalErros(),
		ErrosVazias:    len(r.ErrosVazias),
		ErrosNCM:       len(r.ErrosNCM),
		ErrosCSTOrigem: len(r.ErrosCSTOrigem),
		ErrosCSOSN:     len(r.ErrosCSOSN),
		ErrosTipoItem:  len(r.ErrosTipoItem),
		Detalhes:       detalhes,
	}
}

// MarshalJSON implementa json.Marshaler para garantir campos não-nulos no JSON
func (r RespostaValidacaoAPI) MarshalJSON() ([]byte, error) {
	type Alias RespostaValidacaoAPI
	if r.Detalhes == nil {
		r.Detalhes = []ErroValidacao{}
	}
	return json.Marshal((Alias)(r))
}

// TotalErros retorna a soma de todos os erros
func (r ResultadoValidacaoCompleto) TotalErros() int {
	return len(r.ErrosVazias) + len(r.ErrosNCM) + len(r.ErrosCSTOrigem) + len(r.ErrosCSOSN) + len(r.ErrosTipoItem)
}
