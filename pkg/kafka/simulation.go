package kafka

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)



func (p *Producer) sendSimEvent(event *SimEvent) error {

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: SimEventTopic[SimEventStatus(event.Status)],
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}

func (p *Producer) sendSimRollbackEvent(event *SimEvent) error {

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: SimEventRollbackTopic[SimEventStatus(event.Status)],
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}



func (p *Producer) SendSimRequestedEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusRequested)
	return p.sendSimEvent(event)
}

func (p *Producer) SendSimRequestedRollbackEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusRequested)
	return p.sendSimRollbackEvent(event)
}

func (p *Producer) SendSimPendingEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusPending)
	return p.sendSimEvent(event)
}

func (p *Producer) SendSimPendingRollbackEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusPending)
	return p.sendSimRollbackEvent(event)
}

func (p *Producer) SendSimAllocatedEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusAllocated)
	return p.sendSimEvent(event)
}

func (p *Producer) SendSimAllocatedRollbackEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusAllocated)
	return p.sendSimRollbackEvent(event)
}

func (p *Producer) SendSimFailedEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusFailed)
	return p.sendSimEvent(event)
}

func (p *Producer) SendSimFailedRollbackEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusFailed)
	return p.sendSimRollbackEvent(event)
}

func (p *Producer) SendSimFinishedEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusFinished)
	return p.sendSimEvent(event)
}

func (p *Producer) SendSimFinishedRollbackEvent(event *SimEvent) error {
	event.Status = int64(SimEventStatusFinished)
	return p.sendSimRollbackEvent(event)
}



func (c *Consumer) subscribeSimEvent(status SimEventStatus, callback func(*SimEvent)) error {
	
	pc, err := c.consumer.ConsumePartition(SimEventTopic[status], 0, sarama.OffsetOldest)
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

				callback(event)
			}
		}
	}(c.ctx)

	return nil
}



func (c *Consumer) subscribeSimRollbackEvent(status SimEventStatus, callback func(*SimEvent)) error {
	
	pc, err := c.consumer.ConsumePartition(SimEventRollbackTopic[status], 0, sarama.OffsetOldest)
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

				callback(event)
			}
		}
	}(c.ctx)

	return nil
}



type SimEventListener interface {
	OnReceiveAllSimEvent(*SimEvent)
	OnReceiveAllSimRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeSimRequestedEvent(impl SimEventListener) error {
	if err := c.subscribeSimEvent(SimEventStatusRequested, func(event *SimEvent) {
		impl.OnReceiveAllSimEvent(event)
	}); err != nil {
		return err
	}

	if err := c.subscribeSimRollbackEvent(SimEventStatusRequested, func(event *SimEvent) {
		impl.OnReceiveAllSimRollbackEvent(event)
	}); err != nil {
		return err
	}

	return nil
}





func (c *Consumer) SubscribeSimAllocatedEvent(impl SimEventListener) error {
	return c.subscribeSimEvent(SimEventStatusAllocated, func(event *SimEvent) {
		impl.OnReceiveAllSimEvent(event)
	})
}

func (c *Consumer) SubscribeSimPendingEvent(impl SimEventListener) error {
	return c.subscribeSimEvent(SimEventStatusPending, func(event *SimEvent) {
		impl.OnReceiveAllSimEvent(event)
	})
}


func (c *Consumer) SubscribeSimRollbackEvent(impl SimEventListener) error {

	pc, err := c.consumer.ConsumePartition(simRollbackEventTopicName, 0, sarama.OffsetOldest)
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
	
				impl.OnReceiveAllSimRollbackEvent(event)
			}
		}
	}(c.ctx)

	return nil
}