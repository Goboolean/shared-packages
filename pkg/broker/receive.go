package broker



type SubscribeListener interface {
	OnReceiveMessage(name string, stock *StockAggregate) error
}
