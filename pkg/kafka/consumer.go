package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

type Consumer struct {
	consumer *kafka.Consumer

	topic  string
	config *kafka.ConfigMap

	instance SubscribeListener
}

/*"security.protocol": "SASL_SSL",*/

func NewConsumer(stock string, instance SubscribeListener) *Consumer {

	config := &kafka.ConfigMap{
		"bootstrap.servers":       KAFKA_ADDR,
		"sasl.mechanism":          "PLAIN",
		"sasl.username":           KAFKA_USER,
		"sasl.password":           KAFKA_PASS,
		"auto.offset.reset":       "earliest",
		"socket.keepalive.enable": true,
	}

	consumer, err := kafka.NewConsumer(config)

	if err != nil {
		log.Fatalf("err: failed to laod kafka consumer: %v", err)
		return nil
	}

	if err := consumer.Subscribe(stock, nil); err != nil {
		log.Fatalf("err: failed to subscribe topic: %v", err)
		return nil
	}

	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)

			if err != nil {
				log.Fatalf("err: failed to read received message: %v", err)
			} else {
				data := StockAggregate{}

				if err := proto.Unmarshal(msg.Value, &data); err != nil {
					log.Fatalf("err: failed to deserialize message: %v", err)
				}

				instance.OnReceiveMessage(&data)
			}
		}
	}()

	return &Consumer{
		topic:  stock,
		config: config,
		instance: instance,
	}
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}
