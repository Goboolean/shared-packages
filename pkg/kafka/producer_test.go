package kafka_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/kafka"
	"github.com/Goboolean/shared-packages/pkg/resolver"
)

var pub *kafka.Producer

func SetupProducer() {
	pub = kafka.NewProducer(&resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	})
}

func TeardownProducer() {
	if err := pub.Close(); err != nil {
		panic(err)
	}
}


func TestProducer(t *testing.T) {
	SetupProducer()
	TeardownProducer()
}