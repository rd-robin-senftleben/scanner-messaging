package message

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

type KafkaConsumer struct {
	Consumer *kafka.Consumer
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
		Consumer: c,
	}
}

func (kc KafkaConsumer) Read() *AssetMessage {
	ev := kc.Consumer.Poll(200)

	if ev == nil {
		return nil
	}

	switch e := ev.(type) {
	case *kafka.Message:
		fmt.Printf("%% Message on %s:\n%s\n",
			e.TopicPartition, string(e.Value))
		if e.Headers != nil {
			fmt.Printf("%% Headers: %v\n", e.Headers)
		}
		assetMessage := &AssetMessage{}
		err := json.Unmarshal(e.Value, assetMessage)
		if err != nil {
			return nil
		}

		return assetMessage
	case kafka.Error:
		fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
		if e.Code() == kafka.ErrAllBrokersDown {
			return nil
		}
	default:
		fmt.Printf("Ignored %v\n", e)
	}

	return nil
}
