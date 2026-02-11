package excel

import (
	"Parser2/internal/domain"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Validator valida células do excel
type Validator struct {
	arquivo    *excelize.File
	sheetName  string
	cabecalhos []string
}

// NovoValidator cria instancia de Validator (cria validator)
func NovoValidator(arquivo *excelize.File, sheetName string, cabecalhos []string) *Validator {
	return &Validator{arquivo, sheetName, cabecalhos}
}

// ValidarTudo valida todas as linhas
func (v *Validator) ValidarTudo(totalLinhas int) []domain.ErroValidacao {
	var erros []domain.ErroValidacao

	for linha := 2; linha <= totalLinhas+1; linha++ {
		errosLinha := v.validarLinha(linha)
		erros = append(erros, errosLinha...)
	}
	return erros
}

// validarLinha valida uma linha específica
func (v *Validator) validarLinha(numLinha int) []domain.ErroValidacao {
	var erros []domain.ErroValidacao

	for i := 0; i < len(v.cabecalhos); i++ {
		colName, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			continue
		}

		cellRef := colName + strconv.Itoa(numLinha)
		cellValue, err := v.arquivo.GetCellValue(v.sheetName, cellRef)
		if err != nil {
			continue
		}

		if cellValue == "" {
			erros = append(erros, domain.ErroValidacao{
				Linha:      numLinha,
				Coluna:     colName,
				NomeColuna: v.cabecalhos[i],
				Mensagem:   "está vazio",
			})
		}
	}

	return erros
}
