package main

import (
	"log"
	decoder2 "rssentiallyai/decoder/core"
	"rssentiallyai/kafka/kafka_consumer"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "my_topic"

	consumer, err := kafka_consumer.NewKafkaConsumer(brokers, topic)
	if err != nil {
		log.Fatalln("Failed to start kafka_consumer:", err)
	}

	defer func() {
		if consErr := consumer.Consumer.Close(); consErr != nil {
			log.Fatalln("Failed to close kafka_consumer:", consErr)
		}
	}()

	var decoder decoder2.Decoder = &decoder2.BinaryDecoder{}
	consumer.Consume(decoder)

}
