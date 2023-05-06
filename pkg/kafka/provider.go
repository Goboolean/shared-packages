package kafka



type EventProvider interface {
	OnReceiveSimulationEvent(Event)
	OnReceiveRealTradeEvent(Event)
}

type UnImplementedEventProvider struct {}
func (u *UnImplementedEventProvider) OnReceiveSimulationEvent(Event) {}
func (u *UnImplementedEventProvider) OnReceiveRealTradeEvent(Event) {}