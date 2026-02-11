package main

import (
	"Parser2/cmd"
	"Parser2/internal/config"
	"Parser2/internal/domain"
	"Parser2/internal/excel"
	"Parser2/internal/filesystem"
	"Parser2/internal/formatter"
	"Parser2/logger"
	"fmt"
	"time"
)

func main() {
	// Carregar configurações
	cfg := config.Nova()
	scanner := filesystem.NovoScanner(cfg.CaminhoPadrao)
	arquivos, err := scanner.ListarArquivos()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	//Carrega o menu
	menu := cmd.NovoMenu()
	for {
		resultado := menu.Executar(arquivos)
		switch resultado.Acao {
		case domain.AcaoArquivo:
			processar(resultado.Arquivo.Caminho, cfg)
		case domain.AcaoRefresh:
			arquivos, err = scanner.ListarArquivos()
			if err != nil {
				fmt.Println("Erro:", err)
				return
			}
			fmt.Println("Lista de arquivos atualizada!")
			fmt.Println("Se deseja adicionar um arquivo para leitura, insira-o na pasta: ", cfg.CaminhoPadrao)

		case domain.AcaoSair:
			return
		}
	}
}

// Processar executa a validação completa
func processar(caminho string, cfg *config.Config) {

	// 1. Abrir Excel
	reader, err := excel.NovoReader(caminho, cfg.SheetPadrao)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer reader.Close()

	// 2. Obter metadados
	planilha, err := reader.ObterMetadados()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("\nPlanilha Carregada!\n")
	fmt.Printf("Colunas: %d | Linhas: %d\n", len(planilha.Cabecalhos), planilha.TotalLinhas)
	time.Sleep(2 * time.Second)

	// 3. Validar
	inicio := time.Now()
	validador := excel.NovoValidator(
		reader.ObterArquivo(),
		cfg.SheetPadrao,
		planilha.Cabecalhos,
	)
	erros := validador.ValidarTudo(planilha.TotalLinhas)
	duracao := time.Since(inicio)

	// 4. Processar resultados
	if len(erros) == 0 {
		fmt.Println("\n✓ Nenhuma célula vazia!")
		fmt.Printf("Tempo: %v\n", duracao)
		return
	}

	// 5. Formatar e ordenar
	formatador := formatter.Novo()
	formatador.OrdenarErros(erros)
	mensagens := formatador.FormatarErros(erros)

	// 6. Exibir no console
	fmt.Printf("\n--- ERROS ENCONTRADOS (%d) ---\n", len(erros))
	for _, msg := range mensagens {
		fmt.Println(msg)
	}

	// 7. Salvar log
	caminhoLog, err := logger.SalvarLog(caminho, cfg.DiretorioLogs, mensagens)
	if err != nil {
		fmt.Println("\nErro ao salvar log:", err)
	} else {
		fmt.Printf("\nLog salvo em: %s\n", caminhoLog)
	}

	totalCelulas := planilha.TotalLinhas * len(planilha.Cabecalhos)

	fmt.Printf("Células verificadas: %d\n\n", totalCelulas)
	fmt.Printf("Células vazias: %d\n", len(erros))
	fmt.Printf("Tempo total: %v\n", duracao)
}
