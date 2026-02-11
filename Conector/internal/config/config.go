package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	App      AppConfig
	Firebird FirebirdConfig
	Mongo    MongoConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port int
}
type FirebirdConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
	Charset  string
}
type MongoConfig struct {
	URI string
	DB  string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	//Verificação de variáveis de ambiente OBRIGATÓRIAS
	appPortStr := os.Getenv("APP_PORT")
	if appPortStr == "" {
		return nil, errors.New("APP_PORT é necessaria")
	}
	appPortConv, err := strconv.Atoi(appPortStr)
	if err != nil {
		return nil, errors.New("APP_PORT deve ser um número inteiro")
	}
	cfg.App.Port = appPortConv

	FbHostStr := os.Getenv("FIREBIRD_HOST")
	if FbHostStr == "" {
		return nil, errors.New("FIREBIRD_HOST é necessaria")
	}
	cfg.Firebird.Host = FbHostStr

	FbDBStr := os.Getenv("FIREBIRD_DB")
	if FbDBStr == "" {
		return nil, errors.New("FIREBIRD_DB é necessaria")
	}
	cfg.Firebird.DB = FbDBStr

	FbUserStr := os.Getenv("FIREBIRD_USER")
	if FbUserStr == "" {
		return nil, errors.New("FIREBIRD_USER é necessaria")
	}
	cfg.Firebird.User = FbUserStr

	FbPasswordStr := os.Getenv("FIREBIRD_PASSWORD")
	if FbPasswordStr == "" {
		return nil, errors.New("FIREBIRD_PASSWORD é necessaria")
	}
	cfg.Firebird.Password = FbPasswordStr

	MongoURIStr := os.Getenv("MONGO_URI")
	if MongoURIStr == "" {
		return nil, errors.New("MONGO_URI é necessaria")
	}
	cfg.Mongo.URI = MongoURIStr

	MongoDBStr := os.Getenv("MONGO_DB")
	if MongoDBStr == "" {
		return nil, errors.New("MONGO_DB é necessaria")
	}
	cfg.Mongo.DB = MongoDBStr

	//Verificação de variáveis de ambiente OPCIONAIS, se vazio atribui valor padrão
	appNameStr := os.Getenv("APP_NAME")
	if appNameStr == "" {
		appNameStr = "FireGopher"
	}
	cfg.App.Name = appNameStr

	appEnvStr := os.Getenv("APP_ENV")
	if appEnvStr == "" {
		appEnvStr = "development"
	}
	cfg.App.Env = appEnvStr

	FbPortStr := os.Getenv("FIREBIRD_PORT")
	if FbPortStr == "" {
		FbPortStr = "3050"
	}
	cfg.Firebird.Port, err = strconv.Atoi(FbPortStr)
	if err != nil {
		return nil, errors.New("FIREBIRD_PORT Deve ser um número inteiro")
	}

	FbCharsetStr := os.Getenv("FIREBIRD_CHARSET")
	if FbCharsetStr == "" {
		FbCharsetStr = "UTF8"
	}
	cfg.Firebird.Charset = FbCharsetStr

	return cfg, nil
}
