package broker

import "fmt"



type SubscribeListener interface {
	OnReceiveMessage(name string, stock *StockAggregate)
}



func (c *Subscriber) Subscribe(stock string) error {
	if err := c.Consumer.Subscribe(stock, nil); err != nil {
		return fmt.Errorf("failed to subscribe topic: %v", err)
	}

	return nil
}