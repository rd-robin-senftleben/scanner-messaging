package message

import "github.com/rd-robin-senftleben/scanner-messaging/pkg/message/kafka"

type Consumer interface {
	Read(v any) ([]byte, error)
}

type Producer interface {
	Write(v RequestResponse, topic string)
}

type Messaging struct {
	Consumer Consumer
	Producer Producer
}

func NewMessaging(groupId string) Messaging {
	consumer := kafka.NewConsumer(ALL_TOPICS(), groupId)
	producer := kafka.NewProducer()

	return Messaging{
		Consumer: consumer,
		Producer: producer,
	}
}
