package model

import (
	"context"
	"github.com/cymonevo/cloud-api/internal/database"
	"github.com/cymonevo/cloud-api/internal/mq"

	"github.com/cymonevo/cloud-api/internal/redis"
)

type healthModel struct {
	redisClient redis.Client
	dbClient    database.Client
	publisher   mq.Publisher
}

func (m *healthModel) Call(ctx context.Context) (interface{}, error) {
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	err = m.redisClient.Dial()
	if err != nil {
		return nil, err
	}
	err = m.dbClient.Dial()
	if err != nil {
		return nil, err
	}
	err = m.publisher.Publish("Health", "data")
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *healthModel) Validate(ctx context.Context) error {
	return nil
}
