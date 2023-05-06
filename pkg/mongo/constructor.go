package mongo

import (
	"context"
	"fmt"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	mongo.Client
	DefaultDatabase string
}

func NewDB(c *resolver.Config) *DB {

	if err := c.ShouldUserExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPWExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldHostExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPortExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldDBExist(); err != nil {
		panic(err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority",
		c.User, c.Password, c.Host, c.Port)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	return &DB{
		Client:          *client,
		DefaultDatabase: c.Database,
	}
}
