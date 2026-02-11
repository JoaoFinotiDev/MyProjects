package filesystem

//Scanner lê os arquivos xlsx do diretorio e os retorna de forma ordenada por data de modificação

import (
	"Parser2/internal/domain"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Scanner é uma struct que representa um scanner de arquivos num diretório específico, componente responsável por descobrir
// arquivos xlsx numa pasta.
type Scanner struct {
	Diretorio string
}

// NovoScanner cria uma instancia de Scanner configurada pra varrer o diretório especificado
func NovoScanner(diretorio string) *Scanner {
	return &Scanner{diretorio}
}

// ListarArquivos é um metodo de Scanner que lista todos os arquivos .xlsx no diretório configurado,
// retornando uma slice de domain.ArquivoExcel ordenada por data de modificação (mais recente primeiro).
func (s *Scanner) ListarArquivos() ([]domain.ArquivoExcel, error) {

	//os.ReadDir lê o conteúdo do diretório especificado e retorna uma slice chamada entradas que recebe as entradas de diretório.
	entradas, err := os.ReadDir(s.Diretorio)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar os arquivos: %w", err)
	}

	//Cria uma slice de domain.ArquivoExcel chamada arquivos para armazenar os arquivos encontrados
	var arquivos []domain.ArquivoExcel
	//Percorre cada elemento de entradas
	for _, entrada := range entradas {
		//Ignora diretórios e arquivos que não são .xlsx ou .XLSX
		if entrada.IsDir() || strings.ToLower(filepath.Ext(entrada.Name())) != ".xlsx" {
			continue
		}
		//Obtém métadados do arquivo, principalmente a data de modificação
		infosArquivo, err := entrada.Info()
		if err != nil {
			return nil, fmt.Errorf("erro ao obter info do arquivo '%s': %w", entrada.Name(), err)
		}

		//caminhoCompleto é uma variável que Constrói o caminho completo juntando diretório e nome do arquivo
		caminhoCompleto := filepath.Join(s.Diretorio, entrada.Name())
		//Arquivo recebe uma struct domain.ArquivoExcel contendo nome, caminho e data de modificação
		arquivo := domain.ArquivoExcel{
			Nome:            entrada.Name(),
			Caminho:         caminhoCompleto,
			DataModificacao: infosArquivo.ModTime(),
		}
		//Adiciona o arquivo na slice de arquivos (Arquivos = Arquivos + Arquivo)
		arquivos = append(arquivos, arquivo)
	}
	//Ordena os arquivos por data de modificação, do mais recente para o mais antigo
	sort.Slice(arquivos, func(i, j int) bool {
		return arquivos[i].DataModificacao.After(arquivos[j].DataModificacao)
	})
	//Retorna a slice de arquivos
	return arquivos, nil
}
