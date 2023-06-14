package kafka

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)



type RealEventType int

type RealEventListener interface {
	OnReceiveAllRealEvent(*RealEvent)
}

const (
	RealCreated RealEventType = iota
	RealPending
	RealAllocated
	RealFailed
	RealFinished
)



func (p *Producer) SendRealEvent(event *RealEvent) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: realEventTopicName,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}

func (c *Consumer) SubscribeRealEvent(impl RealEventListener) error {

	pc, err := c.consumer.ConsumePartition(realEventTopicName, 0, sarama.OffsetOldest)
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

				var event *RealEvent
	
				proto.Unmarshal(message.Value, event)
	
				impl.OnReceiveAllRealEvent(event)
			}
		}
	}(c.ctx)

	return nil
}