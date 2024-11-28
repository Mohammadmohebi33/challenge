package mongo

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host   string `koanf:"host"`
	Port   int    `koanf:"port"`
	DBName string `koanf:"db_name"`
}

type MongoDB struct {
	config Config
	client *mongo.Client
	db     *mongo.Database
}

func (m *MongoDB) Conn() *mongo.Database {
	return m.db
}

func New(config Config) (*MongoDB, error) {
	uri := "mongodb://" + config.Host + ":" + strconv.Itoa(config.Port)

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.DBName)
	return &MongoDB{config: config, client: client, db: db}, nil
}

func (m *MongoDB) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return m.client.Disconnect(ctx)
}
