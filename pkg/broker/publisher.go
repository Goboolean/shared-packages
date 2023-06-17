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
	producer *kafka.Producer
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
		"acks": 0,
		"go.delivery.reports": false,
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("failed to create new kafka producer: %v", err)
		return nil
	}

	producer.Events()
	producer.ProduceChannel()

	return &Publisher{producer: producer}
}



func (p *Publisher) Close() {
	p.producer.Close()
}



func (p *Publisher) Ping(ctx context.Context) error {
	deadline, ok := ctx.Deadline()

	if !ok {
		return fmt.Errorf("timeout setting on ctx required")
	}

	remaining := time.Until(deadline)

	_, err := p.producer.GetMetadata(nil, true, int(remaining.Milliseconds()))
	return err
}




func (p *Publisher) SendData(topic string, data *StockAggregate) error {

	topic = packTopic(topic)

	binaryData, err := proto.Marshal(data)

	if err != nil {
		return err
	}

	if err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          binaryData,
	}, nil); err != nil {
		return err
	}

	if remain := p.producer.Flush(int(defaultTimeout.Milliseconds())); remain != 0 {
		return fmt.Errorf("%d events is stil remaining: ", remain)
	}

	return nil
}



func (p *Publisher) SendDataBatch(topic string, batch []*StockAggregate) error {

	topic = packTopic(topic)

	msgChan := p.producer.ProduceChannel()

	binaryBatch := make([][]byte, len(batch))

	for idx := range batch {
		data, err := proto.Marshal(batch[idx])

		if err != nil {
			return err
		}

		binaryBatch[idx] = data
	}

	topic = packTopic(topic)

	for idx := range batch {
		msgChan <- &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          binaryBatch[idx],
		}
	}

	if remain := p.producer.Flush(int(defaultTimeout.Milliseconds())); remain != 0 {
		return fmt.Errorf("%d events is stil remaining: ", remain)
	}

	return nil
}
