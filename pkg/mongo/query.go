package mongo

import (
	"github.com/Goboolean/shared-packages/pkg/resolver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Queries struct {
	client *DB
	tx *mongo.Session
}

func New(db *DB) *Queries {
	return &Queries{client: db}
}



func (q *Queries) InsertStockBatch(tx resolver.Transactioner, stock string, batch []*StockAggregate) error {

	coll := q.client.Database(q.client.DefaultDatabase).Collection(stock)
	session := tx.Transaction().(mongo.Session)

	docs := make([]interface{}, len(batch))

	for idx := range batch {
		docs[idx] = &batch[idx]
	}

	_, err := session.WithTransaction(tx.Context(), func(ctx mongo.SessionContext) (interface{}, error) {
		return coll.InsertMany(ctx, docs)
	})

	return err
}



func (q *Queries) FetchAllStockBatch(tx resolver.Transactioner, stock string) ([]*StockAggregate, error) {
	results := make([]*StockAggregate, 0)

	coll := q.client.Database(q.client.DefaultDatabase).Collection(stock)
	session := tx.Transaction().(mongo.Session)

	_, err := session.WithTransaction(tx.Context(), func(ctx mongo.SessionContext) (interface{}, error) {
		cursor, err := coll.Find(tx.Context(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(tx.Context())	

		for cursor.Next(tx.Context()) {
			var data *StockAggregate
			if err := cursor.Decode(data); err != nil {
				return nil, err
			}

			results = append(results, data)
		}
		return nil, nil		
	})

	return results, err
}



func (q *Queries) FetchAllStockBatchMassive(tx resolver.Transactioner, stock string, stockChan chan *StockAggregate) error {

	coll := q.client.Database(q.client.DefaultDatabase).Collection(stock)
	session := tx.Transaction().(mongo.Session)

	_, err := session.WithTransaction(tx.Context(), func(ctx mongo.SessionContext) (interface{}, error) {

		cursor, err := coll.Find(tx.Context(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(tx.Context())

		for cursor.Next(tx.Context()) {
			var data *StockAggregate
			if err := cursor.Decode(&data); err != nil {
				return nil, err
			}

			stockChan <- data
		}
		return nil, nil
	})

	return err
}