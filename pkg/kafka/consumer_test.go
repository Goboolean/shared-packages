package kafka_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/kafka"
	"github.com/Goboolean/shared-packages/pkg/resolver"
)

var sub *kafka.Consumer

func SetupConsumer() {
	sub = kafka.NewConsumer(&resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	})
}

func TeardownConsumer() {
	if err := sub.Close(); err != nil {
		panic(err)
	}
}


func TestConsumer(t *testing.T) {
	SetupConsumer()
	TeardownConsumer()
}