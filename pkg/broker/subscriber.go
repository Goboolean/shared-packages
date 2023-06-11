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
	OnReceiveStockAggs(name string, stock *StockAggregate)
}

type Subscriber struct {
	consumer *kafka.Consumer

	topic  string
	config *kafka.ConfigMap

	listener SubscribeListener

	ctx context.Context
}



func NewSubscriber(c *resolver.Config, ctx context.Context, lis SubscribeListener) *Subscriber {

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
		consumer: consumer,
		config:     config,
		listener:   lis,
		ctx: ctx,
	}

	pingCtx, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if err := instance.Ping(pingCtx); err != nil {
		panic(err)
	}

	go instance.subscribeMessage(ctx)

	return instance
}



func (s *Subscriber) subscribeMessage(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:

		}

		msg, err := s.consumer.ReadMessage(-1)

		if err != nil {
			log.Fatalf("err: failed to read received message: %v", err)

		} else {
			var data StockAggregate

			if err := proto.Unmarshal(msg.Value, &data); err != nil {
				log.Fatalf("err: failed to deserialize message: %v", err)
			}

			topic := unpackTopic(*msg.TopicPartition.Topic)

			s.listener.OnReceiveStockAggs(topic, &data)
		}
	}
}

func (s *Subscriber) Close() error {
	if err := s.Close(); err != nil {
		return errors.Wrap(err, "failed to subscribe topic")
	}

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

	_, err := s.consumer.GetMetadata(nil, true, int(remaining.Milliseconds()))
	return err
}



func (s *Subscriber) Subscribe(stock string) error {
	if err := s.consumer.Subscribe(stock, nil); err != nil {
		return fmt.Errorf("failed to subscribe topic: %v", err)
	}

	return nil
}