package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SalvarLog cria e salva o arquivo de log com timestamp
// Retorna o caminho completo do arquivo salvo e qualquer erro
func SalvarLog(caminhoArquivoOriginal string, diretorioLogs string, erros []string) (string, error) {
	// 1. Gerar o caminho completo do log
	caminhoLog, err := gerarCaminhoLog(caminhoArquivoOriginal, diretorioLogs)
	if err != nil {
		return "", err
	}

	// 2. Escrever o log no arquivo
	if err := escreverLog(caminhoLog, erros); err != nil {
		return "", err
	}

	return caminhoLog, nil
}

// gerarCaminhoLog cria o caminho completo do log na pasta configurada
func gerarCaminhoLog(caminhoArquivo string, diretorioLogs string) (string, error) {
	// 1. Extrair nome do arquivo sem extensão
	nomeArquivo := filepath.Base(caminhoArquivo)
	semExtensao := strings.TrimSuffix(nomeArquivo, filepath.Ext(nomeArquivo))

	// 2. Gerar nome do log com timestamp
	timestamp := time.Now().Format("20060102_150405")
	nomeLog := fmt.Sprintf("log_validacao_%s_%s.txt", semExtensao, timestamp)

	// 3. Criar caminho completo
	caminhoCompleto := filepath.Join(diretorioLogs, nomeLog)

	// 4. Garantir que a pasta logs existe (cria se não existir)
	if err := os.MkdirAll(diretorioLogs, 0755); err != nil {
		return "", fmt.Errorf("erro ao criar diretório de logs '%s': %w", diretorioLogs, err)
	}

	return caminhoCompleto, nil
}

// escreverLog escreve as mensagens de erro no arquivo
func escreverLog(caminhoCompleto string, erros []string) error {
	f, err := os.Create(caminhoCompleto)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de log: %w", err)
	}
	defer f.Close()

	for _, erro := range erros {
		_, err := f.WriteString(erro + "\n")
		if err != nil {
			return fmt.Errorf("erro ao escrever no log: %w", err)
		}
	}

	return nil
}
