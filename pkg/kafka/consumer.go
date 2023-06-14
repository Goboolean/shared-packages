package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/Shopify/sarama"
)



type Consumer struct {
	consumer sarama.Consumer
	ctx context.Context
}



func NewConsumer(c *resolver.Config, ctx context.Context) *Consumer {
	
	if err := c.ShouldHostExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPortExist(); err != nil {
		panic(err)
	}

	c.Address = fmt.Sprintf("%s:%s", c.Host, c.Port)

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{c.Address}, config)

	if err != nil {
		log.Fatalf("err: failed to laod kafka consumer: %v", err)
		return nil
	}

	return &Consumer{
		consumer: consumer,
		ctx: ctx,
	}

}



func (c *Consumer) Close() error {
	return c.consumer.Close()
}
