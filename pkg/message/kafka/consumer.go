package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

type Consumer struct {
	backend *kafka.Consumer
}

func (kc Consumer) Read(v any) ([]byte, error) {
	ev := kc.backend.Poll(200)

	if ev == nil {
		return nil, errors.New("no data received")
	}

	switch e := ev.(type) {
	case *kafka.Message:
		fmt.Printf("%% Message on %s:\n%s\n",
			e.TopicPartition, string(e.Value))
		if e.Headers != nil {
			fmt.Printf("%% Headers: %v\n", e.Headers)
		}

		err := json.Unmarshal(e.Value, v)
		if err != nil {
			return nil, err
		}

		return e.Value, nil
	case kafka.Error:
		fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
		if e.Code() == kafka.ErrAllBrokersDown {
			return nil, errors.New(e.String())
		}
	default:
		fmt.Printf("Ignored %v\n", e)
	}

	return nil, nil
}

func NewConsumer(topics []string, groupId string) Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics(topics, nil)

	return Consumer{
		backend: c,
	}
}
