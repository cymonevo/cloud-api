package provider

import (
	"sync"

	"github.com/cymonevo/cloud-api/internal/elastic"
)

var (
	esClient     elastic.Client
	syncEsClient sync.Once
)

func GetESClient() elastic.Client {
	if esClient == nil {
		syncEsClient.Do(func() {
			cfg := GetAppConfig().ESConfig
			esClient = elastic.NewESClient(cfg)
		})
	}
	return esClient
}
