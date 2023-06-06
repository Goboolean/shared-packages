package broker_test

import (
	"os"
	"testing"
	"context"
	"reflect"
	"time"

	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/Goboolean/shared-packages/pkg/broker"
)


var (
	sub *broker.Subscriber
)


func SetupSubscriber() {
	sub = broker.NewSubscriber(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	}, &SubscribeListenerImpl{})
}


func TeardownSubscriber() {
	if err := sub.Close(); err != nil {
		panic(err)
	}
}




func TestSubscriber(t *testing.T) {

	SetupSubscriber()
	TeardownSubscriber()
}




type SubscribeListenerImpl struct {}

var stockChan chan *broker.StockAggregate

func (i *SubscribeListenerImpl) OnReceiveStockAggs(name string, data *broker.StockAggregate) {
	stockChan <- data
}


func TestSubscribe(t *testing.T) {

	SetupSubscriber()

	type args struct {
		topic string
		data  *broker.StockAggregate
	}

	tests := []struct {
		name string
		args args
		want *broker.StockAggregate
	}{
		{
			name: "send mock data",
			args: args{
				topic: "mock",
				data: &broker.StockAggregate{
					Average: 1234,
				},
			},
			want: &broker.StockAggregate{
				Average: 1234,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)

			if err := pub.SendData(tt.args.topic, tt.args.data); err != nil {
				t.Errorf("SendData() = %v", err)
			}

			select {
			case <-ctx.Done():
				t.Errorf("timeout: failed to receive data")
			case got := <- stockChan:
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReplaceD() = %v, want %v", got, tt.want)
				}
			}

			cancel()
		})
	}

	if err := pub.SendData(topic, data); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}

	close(stockChan)

	TeardownSubscriber()
}