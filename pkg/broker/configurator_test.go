package broker_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Goboolean/shared-packages/pkg/broker"
	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/joho/godotenv"
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

func TestConfigurator(t *testing.T) {

	conf := broker.NewConfigurator(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	})

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()

	if err := conf.Ping(ctx); err != nil {
		t.Errorf("Ping() = %v", err)
	}

	if err := conf.CreateTopic(ctx, topic); err != nil {
		t.Errorf("CreateTopic() = %v", err)
	}

	if err := conf.DeleteTopic(ctx, topic); err != nil {
		t.Errorf("DeleteTopic() = %v", err)
	}
}
