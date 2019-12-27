package provider

import (
	"sync"

	"github.com/cymonevo/cloud-api/internal/config"
)

var (
	mainConfig     *config.MainConfig
	syncMainConfig sync.Once
)

func GetAppConfig() *config.MainConfig {
	if mainConfig == nil {
		syncMainConfig.Do(func() {
			mainConfig = config.LoadMainConfig()
		})
	}
	return mainConfig
}
