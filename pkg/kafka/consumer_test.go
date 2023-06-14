package kafka_test

import (
	"context"
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/kafka"
	"github.com/Goboolean/shared-packages/pkg/resolver"
)









func TestConsumer(t *testing.T) {

	sub := kafka.NewConsumer(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	}, context.TODO())

	if err := sub.Close(); err != nil {
		t.Errorf("NewConsumer() failed: %v", err)
	}
}