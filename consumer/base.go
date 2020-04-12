package consumer

import "github.com/cymonevo/cloud-api/internal/mq"

type BaseConsumerHandler interface {
	GetConsumers() []mq.BaseConsumer
}
