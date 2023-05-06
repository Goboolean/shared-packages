package kafka_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/Goboolean/shared-packages/pkg/broker"
)

var (
	topic = "topic_test"
	send = &broker.StockAggregate{}
)


func TestPublisher(t *testing.T) {
	pub := broker.NewPublisher(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	})



	if err := pub.Close(); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}
}