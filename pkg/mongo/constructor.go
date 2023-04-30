package mongo

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	MONGO_HOST     = os.Getenv("MONGO_HOST")
	MONGO_PORT     = os.Getenv("MONGO_PORT")
	MONGO_USER     = os.Getenv("MONGO_USER")
	MONGO_PASS     = os.Getenv("MONGO_PASS")
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	MONGO_URI = fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority",
	MONGO_USER, MONGO_PASS, MONGO_HOST, MONGO_PORT)
)

type DB struct {
	mongo.Client
}


func init() {
	
	if _, exist := os.LookupEnv("MONGO_HOST"); !exist {
		log.Fatalf("error: %s enveironment variable required", "MONGO_HOST")
	}

	if _, exist := os.LookupEnv("MONGO_PORT"); !exist {
		log.Fatalf("error: %s enveironment variable required", "MONGO_PORT")
	}

	if _, exist := os.LookupEnv("MONGO_PASS"); !exist {
		log.Fatalf("error: %s enveironment variable required", "MONGO_USER")
	}

	if _, exist := os.LookupEnv("MONGO_HOST"); !exist {
		log.Fatalf("error: %s enveironment variable required", "MONGO_HOST")
	}

	if _, exist := os.LookupEnv("MONGO_DATABASE"); !exist {
		log.Fatalf("error: %s enveironment variable required", "MONGO_DATABASE")
	}

}



func NewDB() *DB {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	return &DB{
		Client: *client,
	}
}

