package config

import (
	"github.com/IBM/sarama"
	"time"
)

func NewConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	config.Consumer.Fetch.Min = 1073741824               // 1GB for high throughput
	config.Consumer.Fetch.Default = 2147483647           // approximately 2GB for high throughput and within int32 limit
	config.Consumer.MaxWaitTime = 100 * time.Millisecond // Reduce this to as low as your application can tolerate

	return config
}
