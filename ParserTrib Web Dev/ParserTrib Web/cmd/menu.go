package cmd

import (
	"ParserTrib/internal/domain"

	"fmt"
)

// Menu representa o menu principal do programa
type Menu struct{}

// NovoMenu cria uma instância de Menu
func NovoMenu() *Menu {
	return &Menu{}
}

// Executar executa o menu com as suas devidas impressões
func (m *Menu) Executar(arquivos []domain.ArquivoExcel) domain.ResultadoMenu {
	var op int
	fmt.Printf("Quantidade de Arquivos disponíveis para processamento: %d\n", len(arquivos))
	if len(arquivos) == 0 {
		fmt.Println("Nenhum arquivo disponível. Por favor, atualize a lista de arquivos.")
	}

	for i, arquivo := range arquivos {
		if i == 0 {
			fmt.Printf("\033[1m%d - %s\033[0m\n", i+1, arquivo.Nome)
		} else {
			fmt.Printf("%d - %s\n", i+1, arquivo.Nome)
		}
	}
	fmt.Println(len(arquivos)+1, " - Atualizar lista de arquivos")
	fmt.Println(len(arquivos)+2, " - Sair")

	//Loop de seleção de opções
	for {
		fmt.Println("Escolha a opção desejada: ")
		fmt.Scanln(&op)
		if op > 0 && op <= len(arquivos) {
			return domain.ResultadoMenu{Acao: domain.AcaoArquivo, Arquivo: arquivos[op-1]}
		}
		if op == len(arquivos)+1 {
			return domain.ResultadoMenu{Acao: domain.AcaoRefresh}
		}
		if op == len(arquivos)+2 {
			return domain.ResultadoMenu{Acao: domain.AcaoSair}
		}
		fmt.Println("Nenhuma opção válida selecionada")
	}
}
