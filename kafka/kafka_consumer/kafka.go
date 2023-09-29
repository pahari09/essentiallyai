package kafka_consumer

import (
	"bytes"
	"github.com/IBM/sarama"
	"log"
	"rssentiallyai/decoder/core"
	config2 "rssentiallyai/kafka/config"
	"runtime"
)

type KafkaConsumer struct {
	Consumer sarama.Consumer
	topic    string
}

func NewKafkaConsumer(brokers []string, topic string) (*KafkaConsumer, error) {
	config := config2.NewConfig()

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{Consumer: consumer, topic: topic}, nil
}

func (kc *KafkaConsumer) Consume(decoder core.Decoder) {
	partitionList, err := kc.Consumer.Partitions(kc.topic)
	if err != nil {
		log.Printf("Error retrieving partition list for topic %s: %s", kc.topic, err)
		return
	}

	messageChannel := make(chan *sarama.ConsumerMessage)

	for _, partition := range partitionList {
		pc, _ := kc.Consumer.ConsumePartition(kc.topic, partition, sarama.OffsetNewest)

		go func(pc sarama.PartitionConsumer) {
			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered in Consume", r)
				}
			}()

			for message := range pc.Messages() {
				messageChannel <- message // send message to channel
			}
			if pcErr := pc.Close(); pcErr != nil {
				log.Printf("Error closing partition kafka_consumer: %s", pcErr)
			}
		}(pc)
	}

	for i := 0; i < runtime.NumCPU(); i++ { // Create as many goroutines as there are CPUs
		go func() {
			for message := range messageChannel {
				packet := bytes.NewBuffer(message.Value)
				//decodedStruct
				_, err := decoder.Decode(packet)
				if err != nil {
					log.Println("Error decoding packet:", err)
				}
			}
		}()
	}
}
