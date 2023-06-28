package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)



// Configurator has a role for making and deleting topic, checking topic exists, and getting topic list.
type Configurator struct {
	AdminClient *kafka.AdminClient
}

// Constructor throws panic when error occurs
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
		//"debug": "security, broker",
	}

	admin, err := kafka.NewAdminClient(config)

	if err != nil {
		panic(err)
	}

	return &Configurator{AdminClient: admin}
}


// It should be called before program ends to free memory
func (c *Configurator) Close() {
	c.AdminClient.Close()
}


// Check if connection to kafka is alive
func (c *Configurator) Ping(ctx context.Context) error {

	// It requires ctx to be deadline set, otherwise it will return error
	// It will return error if there is no response within deadline
	deadline, ok := ctx.Deadline()

	if !ok {
		return fmt.Errorf("timeout setting on ctx required")
	}

	remaining := time.Until(deadline)

	_, err := c.AdminClient.GetMetadata(nil, true, int(remaining.Milliseconds()))
	return err
}


// Create a topic
func (c *Configurator) CreateTopic(ctx context.Context, topic string) error {

	// It returns error when topic already exists
	topic = packTopic(topic)

	exists, err := c.TopicExists(ctx, topic)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	topicInfo := kafka.TopicSpecification{
		Topic: topic,
		NumPartitions: 1,
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


// Delete a topic
func (c *Configurator) DeleteTopic(ctx context.Context, topic string) error {

	// It returns error when the topic does not exist
	topic = packTopic(topic)

	result, err := c.AdminClient.DeleteTopics(ctx, []string{topic})

	if err != nil {
		return errors.Wrap(err, "fatal error while deleting topic")
	}
	
	if err := result[0].Error; err.Code() != kafka.ErrNoError {
		return errors.Wrap(fmt.Errorf(err.String()), "trival error while deleting topic")
	}

	return nil
}


// Check if given topic exists
func (c *Configurator) TopicExists(ctx context.Context, topic string) (bool, error) {

	topic = packTopic(topic)

	deadline, ok := ctx.Deadline()

	if !ok {
		return false, fmt.Errorf("timeout setting on ctx required")
	}

	remaining := time.Until(deadline)

	metadata, err := c.AdminClient.GetMetadata(nil, true, int(remaining.Milliseconds()))
	if err != nil {
		return false, err
	}

	_, exists := metadata.Topics[topic]
	return exists, nil
}


// Get all existing topic list as a string slice
func (c *Configurator) GetTopicList(ctx context.Context) ([]string, error) {

	deadline, ok := ctx.Deadline()

	if !ok {
		return nil, fmt.Errorf("timeout setting on ctx required")
	}

	remaining := time.Until(deadline)

	metadata, err := c.AdminClient.GetMetadata(nil, true, int(remaining.Milliseconds()))
	if err != nil {
		return nil, err
	}

	topicList := make([]string, 0)

	for topic := range metadata.Topics {
		if len(topic) > 0 {
			topicList = append(topicList, topic)
		}
	}

	return topicList, nil
}