package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rd-robin-senftleben/scanner-messaging/pkg/message"
)

type Producer struct {
	backend *kafka.Producer
}

func NewProducer() Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return Producer{
		backend: p,
	}
}

func (kc Producer) Write(v message.RequestResponse, topic string) {
	out, _ := json.Marshal(v)

	kc.backend.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          out,
	}, nil)

	// Wait for message deliveries before shutting down
	kc.backend.Flush(15 * 1000)
}
