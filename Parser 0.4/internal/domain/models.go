package domain

import (
	"fmt"
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

// ResultadoValidacao armazena os resultados da validação
type ResultadoValidacao struct {
	TotalErros    int
	Erros         []string
	tempoExecucao time.Duration
}

// ErroValidacao representa um erro especifico
type ErroValidacao struct {
	Linha      int
	Coluna     string
	NomeColuna string
	Mensagem   string
}

// String formatada de mensagem de erro para exibição
func (e ErroValidacao) String() string {
	return fmt.Sprintf("[ERRO] Linha %d, Coluna %s (%s): CÉLULA VAZIA",
		e.Linha,
		e.Coluna,
		e.NomeColuna)
}
