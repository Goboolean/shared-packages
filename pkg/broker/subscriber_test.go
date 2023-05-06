package broker_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/Goboolean/shared-packages/pkg/broker"
)


var (
	sub *broker.Subscriber
)


func SetupSubscriber() {
	sub = broker.NewSubscriber(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	}, &SubscribeListenerImpl{})
}


func TeardownSubscriber() {
	if err := sub.Close(); err != nil {
		panic(err)
	}
}




func TestSubscriber(t *testing.T) {

	SetupSubscriber()
	TeardownSubscriber()
}