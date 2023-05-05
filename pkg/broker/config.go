package broker

import (
	"fmt"
	"log"
	"os"
)

var (
	KAFKA_HOST = os.Getenv("KAFKA_HOST")
	KAFKA_PORT = os.Getenv("KAFKA_PORT")
	KAFKA_USER = os.Getenv("KAFKA_USER")
	KAFKA_PASS = os.Getenv("KAFKA_PASS")
	KAFKA_ADDR = fmt.Sprintf("%s:%s", KAFKA_HOST, KAFKA_PORT)
)

func init() {
	if _, exist := os.LookupEnv("KAFKA_HOST"); !exist {
		log.Fatalf("error: %s enveironment variable required", "KAFKA_HOST")
	}

	if _, exist := os.LookupEnv("KAFKA_PORT"); !exist {
		log.Fatalf("error: %s enveironment variable required", "KAFKA_PORT")
	}

	if _, exist := os.LookupEnv("KAFKA_USER"); !exist {
		log.Fatalf("error: %s enveironment variable required", "KAFKA_USER")
	}

	if _, exist := os.LookupEnv("KAFKA_PASS"); !exist {
		log.Fatalf("error: %s enveironment variable required", "KAFKA_PASS")
	}
}
