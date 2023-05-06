package broker_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/broker"
	"github.com/Goboolean/shared-packages/pkg/resolver"
)

var (
	pub *broker.Publisher
	data = &broker.StockAggregate{}
	dataBatch = []broker.StockAggregate{
		{}, {}, {},
	}
)



func SetupPublisher() {
	pub = broker.NewPublisher(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	})
}

func TeardownPublisher() {
	if err := pub.Close(); err != nil {
		panic(err)
	}
}



func TestPublisher(t *testing.T) {

	SetupPublisher()
	TeardownPublisher()

}