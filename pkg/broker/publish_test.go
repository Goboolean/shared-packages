package broker_test

import 	"testing"



func TestSendData(t *testing.T) {

	SetupPublisher()

	if err := pub.SendData(topic, data); err != nil {
		t.Errorf("SendData() failed: %v", err)
	}

	TeardownPublisher()
}


func TestSendDataBatch(t *testing.T) {

	SetupPublisher()

	if err := pub.SendDataBatch(topic, dataBatch); err != nil {
		t.Errorf("SendDataBatch() failed: %v", err)
	}

	TeardownPublisher()
}