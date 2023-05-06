package broker

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)



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


