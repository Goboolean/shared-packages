package broker_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/Goboolean/shared-packages/pkg"
	"github.com/Goboolean/shared-packages/pkg/broker"
)

var (
	topic = "topic_test"
	data = &broker.StockAggregate{}
)


func TestMain(m *testing.M) {

	if err := os.Chdir("../../"); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	
	code := m.Run()

	os.Exit(code)
}


func TestPublisher(t *testing.T) {

	pub := broker.NewPublisher(&pkg.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	})

	if err := pub.SendData(topic, data); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}

	if err := pub.Close(); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}
}