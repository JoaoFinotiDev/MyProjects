package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env:", err)
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI não definida no .env")
	}
	testMongoConnection(uri)
}

func testMongoConnection(uri string) {
	fmt.Println("\nTeste de conexão MongoDB")
	fmt.Printf("URI: %s\n", uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Erro ao conectar MongoDB:", err)
	}
	defer client.Disconnect(ctx) // ← Movido para cá (logo após Connect)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Erro no ping MongoDB:", err) // ← Adicionado tratamento
	}

	fmt.Println("✅ CONEXÃO MONGO OK!")
}
