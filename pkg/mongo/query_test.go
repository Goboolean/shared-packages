package mongo_test

import (
	"context"
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg"
	"github.com/Goboolean/shared-packages/pkg/mongo"
	"github.com/joho/godotenv"
)

var (
	instance *mongo.DB
	queries *mongo.Queries
)

var (
	stockName = "asdf"
	stockBatch = []mongo.StockAggregate{
		{},
		{},
	}
)


var testStockName string = ""



func TestMain(m *testing.M) {

	if err := os.Chdir("../../"); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	c := pkg.Config{
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		Password: os.Getenv("MONGO_PASS"),
		Database: os.Getenv("MONGO_DATABASE"),
	}
	instance = mongo.NewDB(&c)
	queries = mongo.New(instance)

	code := m.Run()

	instance.Disconnect(context.TODO())

	os.Exit(code)
}



func TestInsertStockBatch(t *testing.T) {
	session, err := instance.StartSession()
	if err != nil {
		t.Errorf("failed to start session: %v", err)
	}

	tx := mongo.NewTransaction(session, context.TODO())

	if err := queries.InsertStockBatch(*tx, stockName, stockBatch); err != nil {
		t.Errorf("failed to insert: %v", err)
	}
}


func TestFetchAllStockBatch(t *testing.T) {

	session, err := instance.StartSession()
	if err != nil {
		t.Errorf("failed to start session: %v", err)
	}

	tx := mongo.NewTransaction(session, context.TODO())

	stockChan := make(chan mongo.StockAggregate, 100)

	if err := queries.FetchAllStockBatch(*tx, stockName, stockChan); err != nil {
		t.Errorf("failed to fetch: %v", err)
	}

	// check equals
}
