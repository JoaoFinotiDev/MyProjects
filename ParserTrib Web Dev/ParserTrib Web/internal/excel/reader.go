package excel

import (
	"ParserTrib/internal/domain"
	"fmt"
	"github.com/xuri/excelize/v2"
)

// Reader que encapsula operações de leitura do Excel
type Reader struct {
	arquivo   *excelize.File
	sheetName string
}

// NovoReader cria uma instância de Reader (um novo leitor de Excel)
func NovoReader(caminho, sheetName string) (*Reader, error) {
	f, err := excelize.OpenFile(caminho)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	return &Reader{
		arquivo:   f,
		sheetName: sheetName,
	}, nil
}

// Close fecha o arquivo
func (r *Reader) Close() error {
	return r.arquivo.Close()
}

// ObterMetadados retorna informações da planilha
func (r *Reader) ObterMetadados() (*domain.Planilha, error) {
	cabecalhos, err := r.lerCabecalho()
	if err != nil {
		return nil, err
	}

	totalLinhas, err := r.obterTotalLinhas()
	if err != nil {
		return nil, err
	}

	return &domain.Planilha{
		NomeSheet:   r.sheetName,
		Cabecalhos:  cabecalhos,
		TotalLinhas: totalLinhas,
	}, nil
}

// lerCabecalho (privada para uso interno)
func (r *Reader) lerCabecalho() ([]string, error) {
	rows, err := r.arquivo.GetRows(r.sheetName)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("planilha '%s' está vazia", r.sheetName)
	}

	return rows[0], nil
}

// ObterTodasLinhas retorna todas as linhas da planilha
func (r *Reader) ObterTodasLinhas() ([][]string, error) {
	return r.arquivo.GetRows(r.sheetName)
}

// obterTotalLinhas (privada para uso interno)
func (r *Reader) obterTotalLinhas() (int, error) {
	rows, err := r.arquivo.GetRows(r.sheetName)
	if err != nil {
		return 0, err
	}

	if len(rows) <= 1 {
		return 0, nil
	}

	return len(rows) - 1, nil
}

// ObterArquivo retorna arquvio excelize (para validador)
func (r *Reader) ObterArquivo() *excelize.File {
	return r.arquivo
}
