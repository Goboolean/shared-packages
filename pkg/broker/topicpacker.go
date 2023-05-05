package broker

import "strings"


var prefix string = "broker:"

func packTopic(topic string) string {
	return prefix + topic
}

func unpackTopic(topic string) string {
	return strings.TrimPrefix(topic, prefix)
}