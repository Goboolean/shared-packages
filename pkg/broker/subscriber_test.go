package broker_test

import (
	"testing"

	"github.com/Goboolean/shared-packages/pkg/broker"
)


type SubscribeListenerImpl struct {}

func (i *SubscribeListenerImpl) OnReceiveMessage(stock *broker.StockAggregate) error {
	received = stock
	return nil
}

var received *broker.StockAggregate



func TestSubscriber(t *testing.T) {
	sub := broker.NewSubscriber(topic, &SubscribeListenerImpl{})

	if sub == nil {
		t.Errorf("NewSubscriber() failed: see log.Fatal")
	}

	pub := broker.NewPublisher()
	if sub == nil {
		t.Errorf("NewPublisher() failed: see log.Fatal")
	}

	if err := pub.SendData(topic, &broker.StockAggregate{}); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}

	if received != send {
		t.Errorf("received %v, want %v", received, send)
	}

	if err := sub.Close(); err != nil {
		t.Errorf("NewSubscriber() failed: %v", err)
	}
}