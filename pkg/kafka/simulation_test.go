package kafka_test

import (
	"context"
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/kafka"
	"github.com/Goboolean/shared-packages/pkg/resolver"
	"google.golang.org/protobuf/proto"
)


var (
	simEventSend = &kafka.SimEvent{}
	simEventReceivedList = make([]*kafka.SimEvent, 0)
)

type SimEventTester struct {}

func (t *SimEventTester) OnReceiveAllSimEvent(e *kafka.SimEvent) {
	simEventReceivedList = append(simEventReceivedList, e)
}

func (t *SimEventTester) OnReceiveAllSimRollbackEvent(e *kafka.SimEvent) {}


func TestSimulation(t *testing.T) {

	pub := kafka.NewProducer(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),		
	})

	sub := kafka.NewConsumer(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),		
	}, context.TODO())

	if err := sub.SubscribeSimEvent(&SimEventTester{}); err != nil {
		t.Errorf("failed to run SubscribeSimEvent(): %v", err)
	}

	if err := pub.SendSimEvent(simEventSend); err != nil {
		t.Errorf("failed to run SendSimEvent(): %v", err)
	}

	if (func() bool {
		for _, received := range simEventReceivedList {
			if proto.Equal(simEventSend, received) {
				return true
			}
		}
		return false
	}()) {
		t.Errorf("failed to run OnReceiveAllSimEvent()")
	}	
}