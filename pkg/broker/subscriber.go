package broker

import (
	"fmt"
	"log"

	"github.com/Goboolean/shared-packages/pkg"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type Subscriber struct {
	consumer *kafka.Consumer

	topic  string
	config *kafka.ConfigMap

	listener SubscribeListener

	flagClosed chan struct{}
}



func NewSubscriber(c *pkg.Config, lis SubscribeListener) *Subscriber {

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

	flagClosed := make(chan struct{})

	go func(ch chan struct{}) {
		for {
			msg, err := consumer.ReadMessage(-1)

			select {
			case <-ch:
				return
			default:
				
			}

			if err != nil {
				log.Fatalf("err: failed to read received message: %v", err)
				return

			} else {
				var data StockAggregate

				if err := proto.Unmarshal(msg.Value, &data); err != nil {
					log.Fatalf("err: failed to deserialize message: %v", err)
				}
				topic := msg.TopicPartition.Topic

				lis.OnReceiveMessage(*topic, &data)
			}
		}
	}(flagClosed)

	return &Subscriber{
		config: config,
		listener: lis,
		flagClosed: flagClosed,
	}
}

func (c *Subscriber) Close() error {
	if err := c.consumer.Close(); err != nil {
		return errors.Wrap(err, "failed to subscribe topic")
	}

	c.flagClosed <- struct{}{}
	return nil
}

func (c *Subscriber) Subscribe(stock string) error {
	if err := c.consumer.Subscribe(stock, nil); err != nil {
		return fmt.Errorf("failed to subscribe topic: %v", err)
	}

	return nil
}
