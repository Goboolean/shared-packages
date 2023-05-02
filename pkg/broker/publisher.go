package broker

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Publisher struct {
	producer *kafka.Producer
}

/*"security.protocol": "SASL_SSL",*/

func NewPublisher() *Publisher {
	config := &kafka.ConfigMap{
		"bootstrap.servers": KAFKA_ADDR,
		"sasl.mechanism":    "PLAIN",
		"sasl.username":     KAFKA_USER,
		"sasl.password":     KAFKA_PASS,
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("failed to create new kafka producer: %v", err)
		return nil
	}

	return &Publisher{producer: producer}
}

//protoc -I api/proto --go_out=pkg/kafka --go_opt=paths=source_relative stockaggs.proto

func (p *Publisher) Close() error {
	p.producer.Close()
	return nil
}
