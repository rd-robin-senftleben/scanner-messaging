package message

import (
	"github.com/rd-robin-senftleben/scanner-messaging/pkg/message/kafkabackend"
	"github.com/rd-robin-senftleben/scanner-messaging/pkg/message/types"
)

type Consumer interface {
	Read(v any) ([]byte, error)
}

type Producer interface {
	Write(v types.RequestResponse, topic string)
}

type Messaging struct {
	Consumer Consumer
	Producer Producer
}

func NewMessaging(groupId string) Messaging {
	consumer := kafkabackend.NewConsumer(types.ALL_TOPICS(), groupId)
	producer := kafkabackend.NewProducer()

	return Messaging{
		Consumer: consumer,
		Producer: producer,
	}
}
