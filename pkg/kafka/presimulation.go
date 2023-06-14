package kafka

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)



type PreSimEventType int

type PreSimEventListener interface {
	OnReceiveAllPreSimEvent(*SimEvent)
}

const (
	PreSimCreated PreSimEventType = iota
	PreSimPending
	PreSimAllocated
	PreSimFailed
	PreSimFinished
)



func (p *Producer) SendPreSimEvent(event *SimEvent) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: simEventTopicName,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}

func (c *Consumer) SubscribePreSimEvent(impl PreSimEventListener) error {

	pc, err := c.consumer.ConsumePartition(preSimEventTopicName, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	go func(ctx context.Context) {

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			for err := range pc.Errors() {
				log.Fatal(err)
			}

			for message := range pc.Messages() {

				var event *SimEvent
	
				proto.Unmarshal(message.Value, event)
	
				impl.OnReceiveAllPreSimEvent(event)
			}
		}
	}(c.ctx)

	return nil
}