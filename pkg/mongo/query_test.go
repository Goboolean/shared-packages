package mongo_test

import (
	"context"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/mongo"
)



var (
	stockName = "test-stock"
	stockBatch = []*mongo.StockAggregate{
		{},
		{},
	}
)




func TestInsertStockBatch(t *testing.T) {
	session, err := db.StartSession()
	if err != nil {
		t.Errorf("failed to start session: %v", err)
	}

	tx := mongo.NewTransaction(session, context.TODO())

	if err := queries.InsertStockBatch(tx, stockName, stockBatch); err != nil {
		t.Errorf("failed to insert: %v", err)
	}
}


func TestFetchAllStockBatchMassive(t *testing.T) {

	session, err := db.StartSession()
	if err != nil {
		t.Errorf("failed to start session: %v", err)
	}

	tx := mongo.NewTransaction(session, context.TODO())

	stockChan := make(chan *mongo.StockAggregate, 100)

	if err := queries.FetchAllStockBatchMassive(tx, stockName, stockChan); err != nil {
		t.Errorf("failed to fetch: %v", err)
	}

	// check equals
}
