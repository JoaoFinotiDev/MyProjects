package main

import (
	"ParserTrib/cmd"
	"ParserTrib/internal/config"
	"ParserTrib/internal/domain"
	"ParserTrib/internal/excel"
	"ParserTrib/internal/filesystem"
	"ParserTrib/internal/formatter"
	"ParserTrib/logger"
	"fmt"
	"os"
	"time"
)

func main() {
	cfg := config.Nova()

	// Se rodar com argumento "server", sobe a API — senão, modo CLI original
	if len(os.Args) > 1 && os.Args[1] == "server" {
		cmd.IniciarServidor(cfg)
		return
	}

	// Modo CLI (comportamento original)
	scanner := filesystem.NovoScanner(cfg.CaminhoPadrao)
	arquivos, err := scanner.ListarArquivos()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

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
			fmt.Println("\n✓ Lista de arquivos atualizada!")
			fmt.Println("Se deseja adicionar um arquivo para leitura, insira-o na pasta:", cfg.CaminhoPadrao)
			fmt.Println()

		case domain.AcaoSair:
			fmt.Println("\n👋 Encerrando o programa. Até logo!")
			return
		}
	}
}

// processar executa a validação completa com 4 etapas
func processar(caminho string, cfg *config.Config) {
	fmt.Println("\n🔄 Iniciando processamento...")

	reader, err := excel.NovoReader(caminho, cfg.SheetPadrao)
	if err != nil {
		fmt.Println("❌ Erro ao abrir arquivo:", err)
		return
	}
	defer reader.Close()

	planilha, err := reader.ObterMetadados()
	if err != nil {
		fmt.Println("❌ Erro ao ler metadados:", err)
		return
	}

	fmt.Printf("\n✅ Planilha Carregada!\n")
	fmt.Printf("📊 Colunas: %d | Linhas: %d\n", len(planilha.Cabecalhos), planilha.TotalLinhas)

	fmt.Printf("\n📋 Cabeçalhos encontrados:\n")
	for i, cab := range planilha.Cabecalhos {
		if i < 5 || cab == "NCM" || cab == "CST Origem" || cab == "CSOSN" || cab == "Tipo Item" {
			fmt.Printf("   - %s\n", cab)
		}
	}
	if len(planilha.Cabecalhos) > 5 {
		fmt.Printf("   ... e mais %d colunas\n", len(planilha.Cabecalhos)-5)
	}

	fmt.Println("\n⏳ Processando validações...")

	rows, err := reader.ObterTodasLinhas()
	if err != nil {
		fmt.Println("❌ Erro ao ler linhas:", err)
		return
	}

	inicio := time.Now()
	validador := excel.NovoValidator(
		rows,
		cfg.SheetPadrao,
		planilha.Cabecalhos,
	)
	resultado := validador.ValidarTudo(planilha.TotalLinhas)
	duracao := time.Since(inicio)
	resultado.TempoExecucao = duracao

	if resultado.TotalErros() == 0 {
		fmt.Println("\n" + formatarLinha("=", 60))
		fmt.Println("✓ NENHUM ERRO ENCONTRADO!")
		fmt.Println(formatarLinha("=", 60))
		fmt.Printf("✓ Todas as células estão preenchidas\n")
		fmt.Printf("✓ Todos os NCMs estão no formato correto\n")
		fmt.Printf("✓ Todos os CST Origem estão válidos\n")
		fmt.Printf("✓ Todos os CSOSN estão válidos\n")
		fmt.Printf("✓ Todos os Tipo Item estão válidos\n")
		fmt.Printf("\n⏱️  Tempo: %v\n", duracao)
		return
	}

	formatadorErros := formatter.Novo()
	formatadorErros.OrdenarErros(resultado.ErrosVazias)
	formatadorErros.OrdenarErros(resultado.ErrosNCM)
	formatadorErros.OrdenarErros(resultado.ErrosCSTOrigem)
	formatadorErros.OrdenarErros(resultado.ErrosCSOSN)
	formatadorErros.OrdenarErros(resultado.ErrosTipoItem)

	saidaFormatada := formatadorErros.FormatarSaida(resultado)
	fmt.Print(saidaFormatada)

	fmt.Println("\n💾 Salvando log...")
	caminhoLog, err := logger.SalvarLog(caminho, cfg.DiretorioLogs, resultado)
	if err != nil {
		fmt.Println("\n❌ Erro ao salvar log:", err)
	} else {
		fmt.Printf("✅ Log salvo em: %s\n", caminhoLog)
	}

	totalCelulas := planilha.TotalLinhas * len(planilha.Cabecalhos)
	fmt.Println("\n" + formatarLinha("=", 60))
	fmt.Println("📊 ESTATÍSTICAS FINAIS")
	fmt.Println(formatarLinha("=", 60))
	fmt.Printf("📝 Células verificadas: %d\n", totalCelulas)
	fmt.Printf("❌ Células vazias: %d\n", len(resultado.ErrosVazias))
	fmt.Printf("⚠️  Erros NCM: %d\n", len(resultado.ErrosNCM))
	fmt.Printf("⚠️  Erros CST Origem: %d\n", len(resultado.ErrosCSTOrigem))
	fmt.Printf("⚠️  Erros CSOSN: %d\n", len(resultado.ErrosCSOSN))
	fmt.Printf("⚠️  Erros Tipo Item: %d\n", len(resultado.ErrosTipoItem))

	fmt.Printf("🔢 Total de erros: %d\n", resultado.TotalErros())
	fmt.Printf("⏱️  Tempo total: %v\n", duracao)
	fmt.Println(formatarLinha("=", 60))
	fmt.Println()
}

func formatarLinha(char string, tamanho int) string {
	linha := ""
	for i := 0; i < tamanho; i++ {
		linha += char
	}
	return linha
}
