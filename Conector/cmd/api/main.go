package main

import (
	"Conector/internal/config"
	"Conector/internal/infrastructure/database/firebird"
	"Conector/internal/infrastructure/database/mongo"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env:", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Erro ao carregar config:", err)
	}

	log.Printf("Conectando Firebird em %s:%d...", cfg.Firebird.Host, cfg.Firebird.Port)
	fbProvider, err := firebird.NewFirebirdProvider(cfg.Firebird)
	if err != nil {
		log.Fatal("Erro ao conectar Firebird:", err)
	}
	defer fbProvider.Close()

	mgProvider, err := mongo.NewMongoProvider(cfg.Mongo)
	if err != nil {
		log.Fatal("Erro ao conectar MongoDB:", err)
	}
	defer mgProvider.Close()

	log.Println("✅ | Conexão: FirebirdDB")
	log.Println("✅ | Conexão: MongoDB")
}
