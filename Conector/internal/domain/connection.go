package domain

import (
	"errors"
	"time"
)

type Connection struct {
	ClientID   string              `json:"clientId" bson:"clientId"`
	Firebird   FirebirdCredentials `json:"firebird" bson:"firebird"`
	Mongo      MongoCredentials    `json:"mongo" bson:"mongo"`
	SyncConfig SyncConfig          `json:"syncConfig" bson:"syncConfig"`
	Active     bool                `json:"active" bson:"active"`
	CreatedAt  time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt" bson:"updatedAt"`
}

type FirebirdCredentials struct {
	Host     string `json:"host" bson:"host"`
	Port     int    `json:"port" bson:"port"`
	Database string `json:"database" bson:"database"`
	User     string `json:"user" bson:"user"`
	Password string `json:"-" bson:"password"`
	Charset  string `json:"charset" bson:"charset"`
}

type MongoCredentials struct {
	URI      string `json:"-" bson:"uri"`
	Database string `json:"database" bson:"database"`
}

type SyncConfig struct {
	Tabelas    []string `json:"tabelas" bson:"tabelas"`
	Frequencia string   `json:"frequencia" bson:"frequencia"`
}

func (c *Connection) Validate() error {
	// ClientID
	if c.ClientID == "" {
		return errors.New("clientId: Não pode ser vazio")
	}

	// Firebird
	if c.Firebird.Host == "" {
		return errors.New("firebird.host: Não pode ser vazio")
	}
	if c.Firebird.Port <= 0 {
		return errors.New("firebird.port: Deve ser maior que zero")
	}
	if c.Firebird.Database == "" {
		return errors.New("firebird.database: Não pode ser vazio")
	}
	if c.Firebird.User == "" {
		return errors.New("firebird.user: Não pode ser vazio")
	}
	if c.Firebird.Password == "" {
		return errors.New("firebird.password: Não pode ser vazio")
	}
	if c.Firebird.Charset == "" {
		return errors.New("firebird.charset: Não pode ser vazio")
	}

	// Mongo
	if c.Mongo.URI == "" {
		return errors.New("mongo.uri: Não pode ser vazio")
	}
	if c.Mongo.Database == "" {
		return errors.New("mongo.database: Não pode ser vazio")
	}

	// SyncConfig
	if len(c.SyncConfig.Tabelas) == 0 {
		return errors.New("syncConfig.tabelas: Deve ter pelo menos uma tabela")
	}
	if c.SyncConfig.Frequencia == "" {
		return errors.New("syncConfig.frequencia: Não pode ser vazio")
	}

	return nil
}
