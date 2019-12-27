package model

import (
	"context"
	"time"

	base "github.com/cymonevo/cloud-api/internal/base/entity"
	"github.com/cymonevo/cloud-api/internal/base/model"
	"github.com/cymonevo/cloud-api/module/article/entity"
)

type getByIDModel struct {
	model.BaseModel
}

func (m *getByIDModel) Call(ctx context.Context) (interface{}, error) {
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	result := entity.Article{
		Title:       "Golang Project Structure",
		Description: "How to design your golang project structure",
		Timestamp: base.Timestamp{
			CreateBy:   "",
			CreateTime: time.Now(),
		},
	}
	return &result, nil
}

func (m *getByIDModel) Validate(ctx context.Context) error {
	return nil
}
