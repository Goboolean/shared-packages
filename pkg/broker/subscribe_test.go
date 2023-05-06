package broker_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Goboolean/shared-packages/pkg/broker"
)



type SubscribeListenerImpl struct {}

var stockChan chan *broker.StockAggregate

func (i *SubscribeListenerImpl) OnReceiveMessage(name string, data *broker.StockAggregate) {
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