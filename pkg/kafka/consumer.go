package kafka

import (
	"fmt"
	"log"

	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/Shopify/sarama"
)

type Consumer struct {
	consumer sarama.Consumer

	data map[string]chan interface{}
}

func NewConsumer(c *resolver.ConfigMap) *Consumer {

	host, err := c.GetStringKey("HOST")
	if err != nil {
		panic(err)
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%s:%s", host, port)

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{address}, config)

	if err != nil {
		log.Fatalf("err: failed to laod kafka consumer: %v", err)
		return nil
	}

	return &Consumer{
		consumer: consumer,
	}
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}
