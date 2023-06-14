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
	realEventSend = &kafka.RealEvent{}
	realEventReceivedList = make([]*kafka.RealEvent, 0)
)

type RealEventTester struct {}

func (t *RealEventTester) OnReceiveAllRealEvent(e *kafka.RealEvent) {
	realEventReceivedList = append(realEventReceivedList, e)
}


func TestRealtime(t *testing.T) {

	pub := kafka.NewProducer(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),		
	})

	sub := kafka.NewConsumer(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),		
	}, context.TODO())

	if err := sub.SubscribeRealEvent(&RealEventTester{}); err != nil {
		t.Errorf("failed to run SubscribeRealEvent(): %v", err)
	}

	if err := pub.SendRealEvent(realEventSend); err != nil {
		t.Errorf("failed to run SendRealEvent(): %v", err)
	}

	if (func() bool {
		for _, received := range realEventReceivedList {
			if proto.Equal(realEventSend, received) {
				return true
			}
		}
		return false
	}()) {
		t.Errorf("failed to run OnReceiveAllRealEvent()")
	}	
}