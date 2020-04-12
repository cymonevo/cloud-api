package provider

import (
	"github.com/cymonevo/cloud-api/consumer"
	"sync"

	"github.com/cymonevo/cloud-api/internal/mq"
)

var (
	mqPublisher     mq.Publisher
	syncMqPublisher sync.Once

	articleConsumers     []mq.BaseConsumer
	syncArticleConsumers sync.Once
)

func GetPublisher() mq.Publisher {
	syncMqPublisher.Do(func() {
		cfg := GetAppConfig().MQPublisherConfig
		mqPublisher = mq.NewPublisher(cfg)
	})
	return mqPublisher
}

func GetArticleConsumers() []mq.BaseConsumer {
	syncArticleConsumers.Do(func() {
		cfg := GetAppConfig().MQConsumerConfig
		articleConsumers = consumer.NewArticleConsumers(cfg).GetConsumers()
	})
	return articleConsumers
}
