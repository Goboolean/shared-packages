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
	pub.Close()
}



func TestPublisher(t *testing.T) {

	SetupPublisher()
	TeardownPublisher()

}



func TestSendData(t *testing.T) {

	var topic = "test-topic"
	SetupPublisher()

	if err := pub.SendData(topic, data); err != nil {
		t.Errorf("SendData() failed: %v", err)
	}

	TeardownPublisher()
}


func TestSendDataBatch(t *testing.T) {

	var topic = "test-topic"
	SetupPublisher()

	if err := pub.SendDataBatch(topic, dataBatch); err != nil {
		t.Errorf("SendDataBatch() failed: %v", err)
	}

	TeardownPublisher()
}



