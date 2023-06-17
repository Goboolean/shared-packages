package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

type Publisher struct {
	*kafka.Producer
}



func NewPublisher(c *resolver.Config) *Publisher {

	if err := c.ShouldHostExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPortExist(); err != nil {
		panic(err)
	}

	c.Address = fmt.Sprintf("%s:%s", c.Host, c.Port)

	config := &kafka.ConfigMap{
		"bootstrap.servers": c.Address,
		"acks": "1",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("failed to create new kafka producer: %v", err)
		return nil
	}

	return &Publisher{Producer: producer}
}



func (p *Publisher) Close() {
	p.Producer.Close()
}



func (p *Publisher) Ping(ctx context.Context) error {
	deadline, ok := ctx.Deadline()

	if !ok {
		return fmt.Errorf("timeout setting on ctx required")
	}

	remaining := time.Until(deadline)
	if remaining < 0 {
		return fmt.Errorf("timeout")
	}

	metaData, err := p.Producer.GetMetadata(nil, true, int(remaining.Milliseconds()))

	fmt.Println(metaData.OriginatingBroker.Host, metaData.OriginatingBroker.Port, metaData.OriginatingBroker.ID)
	return err
}




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
		data, err := proto.Marshal(&batch[idx])

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
