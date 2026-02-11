package formatter

import (
	"ParserTrib/internal/domain"
	"sort"
	"strconv"
	"strings"
)

// Formatter formata erros para exibição
type Formatter struct{}

// Novo cria instância de Formatter
func Novo() *Formatter {
	return &Formatter{}
}

// FormatarErros converte domain.ErroValidacao para strings formatadas
func (f *Formatter) FormatarErros(erros []domain.ErroValidacao) []string {
	var resultado []string

	for _, erro := range erros {
		resultado = append(resultado, erro.String())
	}

	return resultado
}

// OrdenarErros ordena por coluna (alfabética) e depois linha (numérica)
func (f *Formatter) OrdenarErros(erros []domain.ErroValidacao) {
	sort.Slice(erros, func(i, j int) bool {
		if erros[i].Coluna != erros[j].Coluna {
			return erros[i].Coluna < erros[j].Coluna
		}
		return erros[i].Linha < erros[j].Linha
	})
}

// FormatarSaida formata a saída completa com as 4 seções separadas
func (f *Formatter) FormatarSaida(resultado domain.ResultadoValidacaoCompleto) string {
	var sb strings.Builder

	// SEÇÃO 1: CÉLULAS VAZIAS
	if len(resultado.ErrosVazias) > 0 {
		sb.WriteString("\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")
		sb.WriteString("--- CÉLULAS VAZIAS (")
		sb.WriteString(formatarNumero(len(resultado.ErrosVazias)))
		sb.WriteString(") ---\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")

		for _, erro := range resultado.ErrosVazias {
			sb.WriteString(erro.String())
			sb.WriteString("\n")
		}
	} else {
		sb.WriteString("\n✓ Nenhuma célula vazia encontrada!\n")
	}

	// SEÇÃO 2: ERROS DE VALIDAÇÃO NCM
	if len(resultado.ErrosNCM) > 0 {
		sb.WriteString("\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")
		sb.WriteString("--- ERROS DE VALIDAÇÃO NCM (")
		sb.WriteString(formatarNumero(len(resultado.ErrosNCM)))
		sb.WriteString(") ---\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")

		for _, erro := range resultado.ErrosNCM {
			sb.WriteString(erro.String())
			sb.WriteString("\n")
		}
	}

	// SEÇÃO 3: ERROS DE VALIDAÇÃO CST ORIGEM
	if len(resultado.ErrosCSTOrigem) > 0 {
		sb.WriteString("\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")
		sb.WriteString("--- ERROS DE VALIDAÇÃO CST ORIGEM (")
		sb.WriteString(formatarNumero(len(resultado.ErrosCSTOrigem)))
		sb.WriteString(") ---\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")

		for _, erro := range resultado.ErrosCSTOrigem {
			sb.WriteString(erro.String())
			sb.WriteString("\n")
		}
	}

	// SEÇÃO 4: ERROS DE VALIDAÇÃO CSOSN
	if len(resultado.ErrosCSOSN) > 0 {
		sb.WriteString("\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")
		sb.WriteString("--- ERROS DE VALIDAÇÃO CSOSN (")
		sb.WriteString(formatarNumero(len(resultado.ErrosCSOSN)))
		sb.WriteString(") ---\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")

		for _, erro := range resultado.ErrosCSOSN {
			sb.WriteString(erro.String())
			sb.WriteString("\n")
		}
	}

	// SEÇÃO 5: ERROS DE VALIDAÇÃO TIPO ITEM
	if len(resultado.ErrosTipoItem) > 0 {
		sb.WriteString("\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")
		sb.WriteString("--- ERROS DE VALIDAÇÃO TIPO ITEM (")
		sb.WriteString(formatarNumero(len(resultado.ErrosTipoItem)))
		sb.WriteString(") ---\n")
		sb.WriteString(strings.Repeat("=", 60))
		sb.WriteString("\n")

		for _, erro := range resultado.ErrosTipoItem {
			sb.WriteString(erro.String())
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// formatarNumero converte int para string (helper)
func formatarNumero(n int) string {
	return strconv.Itoa(n)
}
