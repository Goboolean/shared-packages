package broker_test

import (
	"testing"

	"github.com/Goboolean/shared-packages/pkg/broker"
)

var (
	topic = "topic_test"
	send = &broker.StockAggregate{}
)


func TestPublisher(t *testing.T) {
	pub := broker.NewPublisher()

	if pub == nil {
		t.Errorf("NewPublisher() failed: see log.Fatal")
	}

	data := &broker.StockAggregate{}

	if err := pub.SendData(topic, data); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}

	if err := pub.Close(); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}
}