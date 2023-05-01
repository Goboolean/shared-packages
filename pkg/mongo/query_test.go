package mongo_test

import (
	"context"
	"os"
	"github.com/Goboolean/shared-packages/pkg/mongo"
	"testing"
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
	instance = mongo.NewDB()
	queries = mongo.New()

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