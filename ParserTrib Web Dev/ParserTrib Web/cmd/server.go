package cmd

import (
	"ParserTrib/api"
	"ParserTrib/internal/config"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// IniciarServidor sobe o servidor HTTP com Gin
func IniciarServidor(cfg *config.Config) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// CORS — permite requisições do frontend em desenvolvimento
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://127.0.0.1:8080",
			"http://192.168.0.189:8080",
		},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Rotas
	handler := api.NovoHandler(cfg)
	router.POST("/api/validar", handler.ValidarExcel)

	// Health check — útil pra confirmar que o servidor tá rodando
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Porta: usa env PORT se existir, senão 3000
	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "3000"
	}

	fmt.Printf("🚀 Servidor iniciado em http://localhost:%s\n", porta)
	fmt.Printf("📌 Endpoint: POST /api/validar\n")
	fmt.Printf("📌 Health:   GET  /api/health\n\n")

	if err := router.Run(":" + porta); err != nil {
		fmt.Printf("❌ Erro ao iniciar servidor: %v\n", err)
		os.Exit(1)
	}
}
