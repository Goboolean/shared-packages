package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("failed to create new kafka producer: %v", err)
		return nil
	}

	instance := &Publisher{Producer: producer}

	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelFunc()

	if err := instance.Ping(ctx); err != nil {
		panic(err)
	}

	return instance
}



func (p *Publisher) Close() error {
	p.Producer.Close()
	return nil
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