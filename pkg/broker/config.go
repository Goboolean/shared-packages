package broker

import "strings"



// every topic that is related to stock data should be prefixed with this
var prefix string = "broker."

// make topic prefixed
func packTopic(topic string) string {
	return prefix + topic
}

// remove prefix from topic
func unpackTopic(topic string) string {
	return strings.TrimPrefix(topic, prefix)
}