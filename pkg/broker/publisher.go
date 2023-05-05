package broker

import (
	"fmt"
	"log"

	"github.com/Goboolean/shared-packages/pkg"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Publisher struct {
	producer *kafka.Producer
}

/*"security.protocol": "SASL_SSL",*/

func NewPublisher(c *pkg.Config) *Publisher {

	if err := c.ShouldHostExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPortExist(); err != nil {
		panic(err)
	}

	c.Address = fmt.Sprintf("%s:%s", c.Host, c.Port)

	config := &kafka.ConfigMap{
		"bootstrap.servers": c.Address,
		"sasl.mechanism":    "PLAIN",
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
