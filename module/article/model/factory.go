package model

import (
	"github.com/cymonevo/cloud-api/internal/base/model"
	"github.com/cymonevo/cloud-api/internal/database"
	"github.com/cymonevo/cloud-api/internal/mq"
	"github.com/cymonevo/cloud-api/internal/redis"
)

type Factory interface {
	NewGetByIDModel() model.BaseModel
	NewHealthModel() model.BaseModel
}

type articleFactory struct {
	db        database.Client
	redis     redis.Client
	publisher mq.Publisher
}

func NewArticleFactory(db database.Client, redis redis.Client, publisher mq.Publisher) Factory {
	return &articleFactory{
		db:        db,
		redis:     redis,
		publisher: publisher,
	}
}

func (f *articleFactory) NewGetByIDModel() model.BaseModel {
	return &getByIDModel{}
}

func (f *articleFactory) NewHealthModel() model.BaseModel {
	return &healthModel{
		dbClient:    f.db,
		redisClient: f.redis,
		publisher:   f.publisher,
	}
}
