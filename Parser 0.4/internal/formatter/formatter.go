package formatter

import (
	"Parser2/internal/domain"
	"sort"
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
		// Usa o método String() do ErroValidacao
		// (já está formatado corretamente lá)
		resultado = append(resultado, erro.String())
	}

	return resultado
}

// OrdenarErros ordena por coluna (alfabética) e depois linha (numérica)
func (f *Formatter) OrdenarErros(erros []domain.ErroValidacao) {
	sort.Slice(erros, func(i, j int) bool {
		// Se colunas diferentes, ordenar por coluna
		if erros[i].Coluna != erros[j].Coluna {
			return erros[i].Coluna < erros[j].Coluna
		}
		// Se mesma coluna, ordenar por linha
		return erros[i].Linha < erros[j].Linha
	})
}
