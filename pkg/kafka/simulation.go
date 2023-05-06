package kafka

import (
	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)



type SimulationEventType int

const (
	Created SimulationEventType = iota
	Pending
	Allocated
	Failed
	Finished	
)

const simulationEventTopicName = "sim"


func (p *Producer) SendSimulationEvent(event *SimulationEvent) error {

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: simulationEventTopicName,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}

func Subscribe() 