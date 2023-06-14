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
	preSimEventSend = &kafka.PreSimEvent{}
	preSimEventReceivedList = make([]*kafka.PreSimEvent, 0)
)

type PreSimEventTester struct {}

func (t *PreSimEventTester) OnReceiveAllPreSimEvent(e *kafka.PreSimEvent) {
	preSimEventReceivedList = append(preSimEventReceivedList, e)
}


func TestPreSimulation(t *testing.T) {

	pub := kafka.NewProducer(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),		
	})

	sub := kafka.NewConsumer(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),		
	}, context.TODO())

	if err := sub.SubscribePreSimEvent(&PreSimEventTester{}); err != nil {
		t.Errorf("failed to run SubscribePreSimEvent(): %v", err)
	}

	if err := pub.SendPreSimEvent(preSimEventSend); err != nil {
		t.Errorf("failed to run SendPreSimEvent(): %v", err)
	}

	if (func() bool {
		for _, received := range preSimEventReceivedList {
			if proto.Equal(preSimEventSend, received) {
				return true
			}
		}
		return false
	}()) {
		t.Errorf("failed to run OnReceiveAllPreSimEvent()")
	}	
}