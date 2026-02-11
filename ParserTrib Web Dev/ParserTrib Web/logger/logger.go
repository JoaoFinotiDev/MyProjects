package logger

import (
	"ParserTrib/internal/domain"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SalvarLog cria e salva o arquivo de log com timestamp e as 4 seções de erros
func SalvarLog(caminhoArquivoOriginal string, diretorioLogs string, resultado domain.ResultadoValidacaoCompleto) (string, error) {
	caminhoLog, err := gerarCaminhoLog(caminhoArquivoOriginal, diretorioLogs)
	if err != nil {
		return "", err
	}

	if err := escreverLog(caminhoLog, resultado); err != nil {
		return "", err
	}

	return caminhoLog, nil
}

// gerarCaminhoLog cria o caminho completo do log na pasta configurada
func gerarCaminhoLog(caminhoArquivo string, diretorioLogs string) (string, error) {
	nomeArquivo := filepath.Base(caminhoArquivo)
	semExtensao := strings.TrimSuffix(nomeArquivo, filepath.Ext(nomeArquivo))

	timestamp := time.Now().Format("20060102_150405")
	nomeLog := fmt.Sprintf("log_validacao_%s_%s.txt", semExtensao, timestamp)

	caminhoCompleto := filepath.Join(diretorioLogs, nomeLog)

	if err := os.MkdirAll(diretorioLogs, 0755); err != nil {
		return "", fmt.Errorf("erro ao criar diretório de logs '%s': %w", diretorioLogs, err)
	}

	return caminhoCompleto, nil
}

// escreverLog escreve as mensagens de erro no arquivo com 4 seções separadas
func escreverLog(caminhoCompleto string, resultado domain.ResultadoValidacaoCompleto) error {
	f, err := os.Create(caminhoCompleto)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de log: %w", err)
	}
	defer f.Close()

	timestamp := time.Now().Format("02/01/2006 15:04:05")
	f.WriteString(strings.Repeat("=", 80) + "\n")
	f.WriteString(fmt.Sprintf("RELATÓRIO DE VALIDAÇÃO - %s\n", timestamp))
	f.WriteString(strings.Repeat("=", 80) + "\n\n")

	f.WriteString("RESUMO:\n")
	f.WriteString(fmt.Sprintf("- Total de células vazias: %d\n", len(resultado.ErrosVazias)))
	f.WriteString(fmt.Sprintf("- Total de erros NCM: %d\n", len(resultado.ErrosNCM)))
	f.WriteString(fmt.Sprintf("- Total de erros CST Origem: %d\n", len(resultado.ErrosCSTOrigem)))
	f.WriteString(fmt.Sprintf("- Total de erros CSOSN: %d\n", len(resultado.ErrosCSOSN)))
	f.WriteString(fmt.Sprintf("- Total de erros Tipo Item: %d\n", len(resultado.ErrosTipoItem)))

	f.WriteString(fmt.Sprintf("- Total geral de erros: %d\n", resultado.TotalErros()))
	f.WriteString(fmt.Sprintf("- Tempo de execução: %v\n", resultado.TempoExecucao))
	f.WriteString("\n")

	if len(resultado.ErrosVazias) > 0 {
		f.WriteString(strings.Repeat("=", 80) + "\n")
		f.WriteString(fmt.Sprintf("CÉLULAS VAZIAS (%d)\n", len(resultado.ErrosVazias)))
		f.WriteString(strings.Repeat("=", 80) + "\n")

		for _, erro := range resultado.ErrosVazias {
			_, err := f.WriteString(erro.String() + "\n")
			if err != nil {
				return fmt.Errorf("erro ao escrever no log: %w", err)
			}
		}
		f.WriteString("\n")
	}

	if len(resultado.ErrosNCM) > 0 {
		f.WriteString(strings.Repeat("=", 80) + "\n")
		f.WriteString(fmt.Sprintf("ERROS DE VALIDAÇÃO NCM (%d)\n", len(resultado.ErrosNCM)))
		f.WriteString(strings.Repeat("=", 80) + "\n")

		for _, erro := range resultado.ErrosNCM {
			_, err := f.WriteString(erro.String() + "\n")
			if err != nil {
				return fmt.Errorf("erro ao escrever no log: %w", err)
			}
		}
		f.WriteString("\n")
	}

	if len(resultado.ErrosCSTOrigem) > 0 {
		f.WriteString(strings.Repeat("=", 80) + "\n")
		f.WriteString(fmt.Sprintf("ERROS DE VALIDAÇÃO CST ORIGEM (%d)\n", len(resultado.ErrosCSTOrigem)))
		f.WriteString(strings.Repeat("=", 80) + "\n")

		for _, erro := range resultado.ErrosCSTOrigem {
			_, err := f.WriteString(erro.String() + "\n")
			if err != nil {
				return fmt.Errorf("erro ao escrever no log: %w", err)
			}
		}
		f.WriteString("\n")
	}

	if len(resultado.ErrosCSOSN) > 0 {
		f.WriteString(strings.Repeat("=", 80) + "\n")
		f.WriteString(fmt.Sprintf("ERROS DE VALIDAÇÃO CSOSN (%d)\n", len(resultado.ErrosCSOSN)))
		f.WriteString(strings.Repeat("=", 80) + "\n")

		for _, erro := range resultado.ErrosCSOSN {
			_, err := f.WriteString(erro.String() + "\n")
			if err != nil {
				return fmt.Errorf("erro ao escrever no log: %w", err)
			}
		}
		f.WriteString("\n")
	}

	if len(resultado.ErrosTipoItem) > 0 {
		f.WriteString(strings.Repeat("=", 80) + "\n")
		f.WriteString(fmt.Sprintf("ERROS DE VALIDAÇÃO TIPO ITEM (%d)\n", len(resultado.ErrosTipoItem)))
		f.WriteString(strings.Repeat("=", 80) + "\n")

		for _, erro := range resultado.ErrosTipoItem {
			_, err := f.WriteString(erro.String() + "\n")
			if err != nil {
				return fmt.Errorf("erro ao escrever no log: %w", err)
			}
		}
		f.WriteString("\n")
	}
	
	// Rodapé

	f.WriteString(strings.Repeat("=", 80) + "\n")
	f.WriteString("FIM DO RELATÓRIO\n")
	f.WriteString(strings.Repeat("=", 80) + "\n")

	return nil
}
