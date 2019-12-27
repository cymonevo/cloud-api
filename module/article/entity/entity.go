package entity

import (
	"github.com/cymonevo/cloud-api/internal/base/entity"
)

type Article struct {
	Title       string
	Description string
	Content     string
	entity.Timestamp
}
