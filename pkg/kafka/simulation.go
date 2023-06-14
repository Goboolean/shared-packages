package kafka

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)



type SimEventType int

type SimEventListener interface {
	OnReceiveAllSimEvent(*SimEvent)
}

const (
	SimCreated SimEventType = iota
	SimPending
	SimAllocated
	SimFailed
	SimFinished
)



func (p *Producer) SendSimEvent(event *SimEvent) error {

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

func (c *Consumer) SubscribeSimEvent(impl SimEventListener) error {

	pc, err := c.consumer.ConsumePartition(simEventTopicName, 0, sarama.OffsetOldest)
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
	
				impl.OnReceiveAllSimEvent(event)
			}
		}
	}(c.ctx)

	return nil
}