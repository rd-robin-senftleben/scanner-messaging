package scanner

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewConsumer(topics []string, groupId string) KafkaConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics(topics, nil)

	return KafkaConsumer{
		consumer: c,
	}
}

func (kc KafkaConsumer) Read() *AssetMessage {
	msg, err := kc.consumer.ReadMessage(time.Second)

	if err == nil {
		assetMessage := &AssetMessage{}
		err := json.Unmarshal(msg.Value, assetMessage)
		if err == nil {
			return assetMessage
		}
	}

	return nil
}
