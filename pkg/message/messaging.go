package message

import (
	"github.com/rd-robin-senftleben/scanner-messaging/pkg/message/kafkabackend"
	"github.com/rd-robin-senftleben/scanner-messaging/pkg/message/types"
)

type Consumer interface {
	Read(v any) ([]byte, error)
}

type Producer interface {
	Write(v any, topic string)
}

type Messaging struct {
	Consumer Consumer
	Producer Producer
}

func NewMessaging(groupId string, bootstrapServer string) Messaging {
	consumer := kafkabackend.NewConsumer(types.ALL_TOPICS(), groupId, bootstrapServer)
	producer := kafkabackend.NewProducer(bootstrapServer)

	return Messaging{
		Consumer: consumer,
		Producer: producer,
	}
}
