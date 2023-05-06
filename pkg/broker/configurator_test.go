package broker_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/broker"
	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/joho/godotenv"
)

var (
	topic = "test"
	conf *broker.Configurator
)



func TestMain(m *testing.M) {

	if err := os.Chdir("../../"); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}



func SetupConfigurator() {
	conf = broker.NewConfigurator(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	})
}

func TeardownConfigurator() {
	if err := conf.Close(); err != nil {
		panic(err)
	}
}



func TestConfigurator(t *testing.T) {

	//defer func() {
	//	if r := recover(); r != nil {
	//		t.Errorf("Configurator Failed: %v", r)
	//	}
	//}()

	SetupConfigurator()
	TeardownConfigurator()
}


