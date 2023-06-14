package kafka_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/kafka"
	"github.com/Goboolean/shared-packages/pkg/resolver"
)


func TestProducer(t *testing.T) {

	pub := kafka.NewProducer(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	})

	if err := pub.Close(); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}
}