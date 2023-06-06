package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
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



func (c *Configurator) CreateTopic(ctx context.Context, topic string) error {

	topicInfo := kafka.TopicSpecification{
		Topic: topic,
		NumPartitions: 10,
		ReplicationFactor: 1,
	}

	result, err := c.AdminClient.CreateTopics(ctx, []kafka.TopicSpecification{topicInfo})

	if err != nil {
		return err
	}
	
	if err := result[0].Error; err.Code() != kafka.ErrNoError {
		return fmt.Errorf(err.String())
	}

	return nil
}


func (c *Configurator) DeleteTopic(ctx context.Context, topic string) error {

	result, err := c.AdminClient.DeleteTopics(ctx, []string{topic})

	if err != nil {
		return errors.Wrap(err, "fatal error while deleting topic")
	}
	
	if err := result[0].Error; err.Code() != kafka.ErrNoError {
		return errors.Wrap(fmt.Errorf(err.String()), "trival error while deleting topic")
	}


	return nil
}