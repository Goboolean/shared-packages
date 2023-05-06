package broker



type SubscribeListener interface {
	OnReceiveMessage(stock *StockAggregate) error
}
