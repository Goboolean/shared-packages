package broker_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Goboolean/shared-packages/pkg/broker"
	"github.com/Goboolean/shared-packages/pkg/resolver"
)

var conf *broker.Configurator



func SetupConfigurator() {
	conf = broker.NewConfigurator(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	})
}

func TeardownConfigurator() {
	conf.Close()
}



func TestConfigurator(t *testing.T) {
	SetupConfigurator()
	TeardownConfigurator()
}



func TestCreateTopic(t *testing.T) {
	
	var topic = "test-topic"
	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	if err := conf.CreateTopic(ctx, topic); err != nil {
		t.Errorf("CreateTopic() = %v", err)
	}

	exists, err := conf.TopicExists(ctx, topic)
	if err != nil {
		t.Errorf("TopicExists() = %v", err)
	}
	if !exists {
		t.Errorf("TopicExists() = %v, expected = true", exists)
	}
	
	TeardownConfigurator()
}



func TestDeleteTopic(t *testing.T) {

	var topic = "test-topic"
	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	
	if err := conf.DeleteTopic(ctx, topic); err != nil {
		t.Errorf("DeleteTopic() = %v", err)
	}

	exists, err := conf.TopicExists(ctx, topic)
	if err != nil {
		t.Errorf("TopicExists() = %v", err)
	}
	if exists {
		t.Errorf("TopicExists() = %v, expected = false", exists)
	}

	TeardownConfigurator()
}