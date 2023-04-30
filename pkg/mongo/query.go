package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Queries struct {
	client *DB
	tx *mongo.Session
}

func New() *Queries {
	return &Queries{client: NewDB()}
}



func (q *Queries) InsertStockBatch(tx Transaction, stock string, batch []StockAggregate) error {

	coll := q.client.Database(MONGO_DATABASE).Collection(stock)
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



func (q *Queries) FetchAllStockBatch(tx Transaction, stock string, stockChan chan StockAggregate) error {

	coll := q.client.Database(MONGO_DATABASE).Collection(stock)
	session := tx.Transaction().(mongo.Session)

	_, err := session.WithTransaction(tx.Context(), func(ctx mongo.SessionContext) (interface{}, error) {

		cursor, err := coll.Find(tx.Context(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(tx.Context())

		for cursor.Next(tx.Context()) {
			var data StockAggregate
			if err := cursor.Decode(&data); err != nil {
				return nil, err
			}

			stockChan <- data
		}
		return nil, nil

	})

	return err
}

