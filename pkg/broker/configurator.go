package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Configurator struct {
	AdminClient *kafka.AdminClient
}

func NewConfigurator(c *resolver.Config) *Configurator {

	if err := c.ShouldHostExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPortExist(); err != nil {
		panic(err)
	}

	c.Address = fmt.Sprintf("%s:%s", c.Host, c.Port)

	config := &kafka.ConfigMap{
		"bootstrap.servers": c.Address,
	}

	admin, err := kafka.NewAdminClient(config)

	if err != nil {
		panic(err)
	}

	instance := &Configurator{AdminClient: admin}

	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelFunc()

	if err := instance.Ping(ctx); err != nil {
		panic(err)
	}

	return instance
}

func (c *Configurator) Close() error {
	return c.Close()
}

func (c *Configurator) Ping(ctx context.Context) error {
	deadline, ok := ctx.Deadline()

	if !ok {
		return fmt.Errorf("timeout setting on ctx required")
	}

	remaining := time.Until(deadline)
	if remaining < 0 {
		return fmt.Errorf("timeout")
	}

	_, err := c.AdminClient.GetMetadata(nil, true, int(remaining.Milliseconds()))
	return err
}
