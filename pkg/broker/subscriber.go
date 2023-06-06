package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)



type SubscribeListener interface {
	OnReceiveMessage(name string, stock *StockAggregate)
}

type Subscriber struct {
	*kafka.Consumer

	topic  string
	config *kafka.ConfigMap

	listener SubscribeListener

	flagClosed chan struct{}
}



func NewSubscriber(c *resolver.Config, lis SubscribeListener) *Subscriber {

	var (
		flagClosed chan struct{}
	)

	if err := c.ShouldHostExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPortExist(); err != nil {
		panic(err)
	}

	c.Address = fmt.Sprintf("%s:%s", c.Host, c.Port)

	config := &kafka.ConfigMap{
		"bootstrap.servers":       c.Address,
		"sasl.mechanism":          "PLAIN",
		"auto.offset.reset":       "earliest",
		"socket.keepalive.enable": true,
	}

	consumer, err := kafka.NewConsumer(config)

	if err != nil {
		log.Fatalf("err: failed to laod kafka consumer: %v", err)
		return nil
	}

	instance := &Subscriber{
		config:     config,
		listener:   lis,
		flagClosed: flagClosed,
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelFunc()

	if err := instance.Ping(ctx); err != nil {
		panic(err)
	}

	flagClosed = make(chan struct{})

	go func(ch chan struct{}) {
		for {
			select {
			case <-ch:
				return
			default:

			}

			msg, err := consumer.ReadMessage(-1)

			if err != nil {
				log.Fatalf("err: failed to read received message: %v", err)

			} else {
				var data StockAggregate

				if err := proto.Unmarshal(msg.Value, &data); err != nil {
					log.Fatalf("err: failed to deserialize message: %v", err)
				}

				topic := unpackTopic(*msg.TopicPartition.Topic)

				lis.OnReceiveMessage(topic, &data)
			}
		}
	}(flagClosed)

	return instance
}

func (c *Subscriber) Close() error {
	if err := c.Consumer.Close(); err != nil {
		return errors.Wrap(err, "failed to subscribe topic")
	}

	c.flagClosed <- struct{}{}
	return nil
}



func (s *Subscriber) Ping(ctx context.Context) error {
	deadline, ok := ctx.Deadline()

	if !ok {
		return fmt.Errorf("timeout setting on ctx required")
	}

	remaining := time.Until(deadline)
	if remaining < 0 {
		return fmt.Errorf("timeout")
	}

	_, err := s.GetMetadata(nil, true, int(remaining.Milliseconds()))
	return err
}



func (c *Subscriber) Subscribe(stock string) error {
	if err := c.Consumer.Subscribe(stock, nil); err != nil {
		return fmt.Errorf("failed to subscribe topic: %v", err)
	}

	return nil
}