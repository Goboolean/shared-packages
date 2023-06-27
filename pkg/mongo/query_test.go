package mongo_test

import (
	"context"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/mongo"
)



var (
	stockName = "asdf"
	stockBatch = []*mongo.StockAggregate{
		{},
		{},
	}
)


var testStockName string = ""




func TestInsertStockBatch(t *testing.T) {
	tx, err := instance.NewTx(context.Background())
	if err != nil {
		t.Errorf("failed to start transaction: %v", err)
	}

	if err := queries.InsertStockBatch(tx, stockName, stockBatch); err != nil {
		t.Errorf("failed to insert: %v", err)
	}

	if err := tx.Commit(); err != nil {
		t.Errorf("failed to commit transaction: %v", err)
	}
}

func isEqual(send, received []*mongo.StockAggregate) bool {
	if len(send) != len(received) {
		return false
	}
	for idx := range send {
		if send[idx] != received[idx] {
			return false
		}
	}
	return true
}


func TestFetchAllStockBatch(t *testing.T) {

	tx, err := instance.NewTx(context.Background())
	if err != nil {
		t.Errorf("failed to start transaction: %v", err)
	}

	result, err := queries.FetchAllStockBatch(tx, stockName);
	if err != nil {
		t.Errorf("FetchAllStockBatch() failed: %v", err)
	}

	if err := tx.Commit(); err != nil {
		t.Errorf("failed to commit transaction: %v", err)
	}

	if !isEqual(stockBatch, result) {
		t.Errorf("FetchAllStockBatch() failed: send and received is not equal")
	}
}

func FetchAllStockBatchMassive(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := instance.NewTx(ctx)
	if err != nil {
		t.Errorf("failed to start transaction: %v", err)
	}

	stockChan := make(chan *mongo.StockAggregate)

	if err := queries.FetchAllStockBatchMassive(tx, stockName, stockChan); err != nil {
		t.Errorf("FetchAllStockBatchMassive() failed: %v", err)
	}

	received := make([]*mongo.StockAggregate, 0)

	loop:
	for {
		select {
		case <-ctx.Done():
			t.Errorf("FetchAllStockBatchMassive() failed with timeout")
			break loop
		case stock := <-stockChan:
			if stock == nil {
				break loop
			}
			received = append(received, stock)			
		}
	}

	if isEqual(stockBatch, received) {
		t.Errorf("FetchAllStockBatchMassive() failed: send and received are not equal")
	}
}
