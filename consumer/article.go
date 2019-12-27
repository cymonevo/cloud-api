package consumer

import (
	"encoding/json"

	"github.com/cymonevo/cloud-api/internal/config"
	"github.com/cymonevo/cloud-api/internal/log"
	"github.com/cymonevo/cloud-api/internal/mq"
	"github.com/cymonevo/cloud-api/module/article/entity"
	"github.com/nsqio/go-nsq"
)

type articleConsumersImpl struct {
	cfg config.MQConsumerConfig
}

func NewArticleConsumers(cfg config.MQConsumerConfig) BaseConsumerHandler {
	return &articleConsumersImpl{
		cfg: cfg,
	}
}

func (c *articleConsumersImpl) GetConsumers() []mq.BaseConsumer {
	return []mq.BaseConsumer{
		mq.NewConsumer(c.cfg, "Health", "health_check", c.insert),
	}
}

func (c *articleConsumersImpl) insert(msg *nsq.Message) error {
	var data entity.Article
	err := json.Unmarshal(msg.Body, &data)
	if err != nil {
		log.ErrorDetail(log.TagMQ, "error unmarshal message", err)
		return err
	}
	log.InfoDetail(log.TagMQ, "consume data %v", data)
	return nil
}
