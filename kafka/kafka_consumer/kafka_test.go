package kafka_consumer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKafkaConsumer(t *testing.T) {
	brokers := []string{"localhost:9092"}
	topic := "my_topic"

	consumer, err := NewKafkaConsumer(brokers, topic)

	assert.NoError(t, err)
	assert.NotNil(t, consumer)
}
