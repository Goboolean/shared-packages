package kafka

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}



func NewProducer(c *resolver.Config) *Producer {

	if err := c.ShouldHostExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPortExist(); err != nil {
		panic(err)
	}

	c.Address = fmt.Sprintf("%s:%s", c.Host, c.Port)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Transaction.ID = createTransactionID()


	producer, err := sarama.NewSyncProducer([]string{c.Address}, config)

	if err != nil {
		panic(err)
	}

	return &Producer{producer: producer}
}



func createTransactionID() string {
	pid := os.Getpid()
	pidString := strconv.Itoa(pid)
	return pidString
}



func (p *Producer) Close() error {
	return p.producer.Close()
}
