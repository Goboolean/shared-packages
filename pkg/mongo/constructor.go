package mongo

import (
	"context"
	"fmt"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
	DefaultDatabase string
}

func NewDB(c *resolver.ConfigMap) *DB {

	user, err := c.GetStringKey("USER")
	if err != nil {
		panic(err)
	}

	password, err := c.GetStringKey("PASSWORD")
	if err != nil {
		panic(err)
	}

	host, err := c.GetStringKey("HOST")
	if err != nil {
		panic(err)
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		panic(err)
	}

	database, err := c.GetStringKey("DATABASE")
	if err != nil {
		panic(err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority",
		user, password, host, port)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	return &DB{
		client:          client,
		DefaultDatabase: database,
	}
}

func (db *DB) NewTx(ctx context.Context) (resolver.Transactioner, error) {
	session, err := db.client.StartSession()
	if err != nil {
		return nil, err
	}
	return NewTransaction(session, ctx), nil
}

func (db *DB) Close() error {
	return db.client.Disconnect(context.Background());
}

func (db *DB) Ping() error {
	return db.client.Ping(context.Background(), nil)
}