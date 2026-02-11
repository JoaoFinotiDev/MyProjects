package mongo

import (
	"Conector/internal/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoProvider struct {
	client *mongo.Client
}

func NewMongoProvider(cfg config.MongoConfig) (*MongoProvider, error) {
	ConnectString := cfg.URI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(ConnectString))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return &MongoProvider{client}, nil
}

func (p *MongoProvider) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return p.client.Disconnect(ctx)
}
