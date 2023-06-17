package broker

import (
	"strings"
	"time"
)



var defaultTimeout = time.Second * 3



var prefix string = "broker:"

func packTopic(topic string) string {
	return prefix + topic
}

func unpackTopic(topic string) string {
	return strings.TrimPrefix(topic, prefix)
}