package broker

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
)

func (p *Publisher) SendData(topic string, data *StockAggregate) error {

	bsonData, err := proto.Marshal(data)

	if err != nil {
		return err
	}

	if err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          bsonData,
	}, nil); err != nil {
		return err
	}

	if remain := p.Flush(int(defaultTimeout.Milliseconds())); remain != 0 {
		return fmt.Errorf("%d events is stil remaining: ", remain)
	}

	return nil
}

func (p *Publisher) SendDataBatch(topic string, batch []StockAggregate) error {

	msgChan := p.ProduceChannel()

	bsonBatch := make([][]byte, len(batch))

	for idx := range batch {
		data, err := bson.Marshal(&batch[idx])

		if err != nil {
			return err
		}

		bsonBatch[idx] = data
	}

	topic = packTopic(topic)

	for idx := range batch {
		msgChan <- &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          bsonBatch[idx],
		}
	}

	if remain := p.Flush(int(defaultTimeout.Milliseconds())); remain != 0 {
		return fmt.Errorf("%d events is stil remaining: ", remain)
	}

	return nil
}
