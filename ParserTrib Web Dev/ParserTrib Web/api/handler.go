package api

import (
	"ParserTrib/internal/config"
	"ParserTrib/internal/excel"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler encapsula as dependências necessárias para os endpoints
type Handler struct {
	cfg *config.Config
}

// NovoHandler cria uma instância do handler com as configurações
func NovoHandler(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

// ValidarExcel é o endpoint POST /api/validar
// Recebe um arquivo .xlsx via multipart/form-data e retorna os erros de validação
func (h *Handler) ValidarExcel(c *gin.Context) {
	// 1. Receber o arquivo do upload
	arquivo, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro":    "Arquivo não fornecido ou erro no upload",
			"detalhes": err.Error(),
		})
		return
	}
	defer arquivo.Close()

	// 2. Validar extensão
	nomeArquivo := header.Filename
	if !strings.HasSuffix(strings.ToLower(nomeArquivo), ".xlsx") {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Apenas arquivos .xlsx são aceitos",
		})
		return
	}

	// 3. Salvar arquivo temporariamente
	tmpDir, err := os.MkdirTemp("", "parsertrib-upload-*")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao criar diretório temporário",
		})
		return
	}
	defer os.RemoveAll(tmpDir) // limpa após processar

	caminhoTmp := filepath.Join(tmpDir, nomeArquivo)
	arquivoTmp, err := os.Create(caminhoTmp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao criar arquivo temporário",
		})
		return
	}

	buf := make([]byte, 1024*1024) // 1MB por vez
	for {
		n, readErr := arquivo.Read(buf)
		if n > 0 {
			if _, writeErr := arquivoTmp.Write(buf[:n]); writeErr != nil {
				arquivoTmp.Close()
				c.JSON(http.StatusInternalServerError, gin.H{
					"erro": "Erro ao salvar arquivo temporário",
				})
				return
			}
		}
		if readErr != nil {
			break
		}
	}
	arquivoTmp.Close()

	// 4. Abrir com o Reader existente
	reader, err := excel.NovoReader(caminhoTmp, h.cfg.SheetPadrao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro":    fmt.Sprintf("Erro ao abrir arquivo: %v", err),
			"detalhes": "Verifique se o arquivo é um .xlsx válido com a aba '" + h.cfg.SheetPadrao + "'",
		})
		return
	}
	defer reader.Close()

	// 5. Obter metadados
	planilha, err := reader.ObterMetadados()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": fmt.Sprintf("Erro ao ler metadados: %v", err),
		})
		return
	}

	// 6. Obter linhas e validar
	rows, err := reader.ObterTodasLinhas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": fmt.Sprintf("Erro ao ler dados: %v", err),
		})
		return
	}

	inicio := time.Now()
	validador := excel.NovoValidator(rows, h.cfg.SheetPadrao, planilha.Cabecalhos)
	resultado := validador.ValidarTudo(planilha.TotalLinhas)
	resultado.TempoExecucao = time.Since(inicio)
	resultado.NomeArquivo = nomeArquivo

	// 7. Converter para resposta da API e retornar
	resposta := resultado.ToRespostaAPI()
	c.JSON(http.StatusOK, resposta)
}
