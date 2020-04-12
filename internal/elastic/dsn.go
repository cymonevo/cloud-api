package elastic

import (
	"fmt"

	"github.com/cymonevo/cloud-api/internal/config"
)

func parseAddress(cfg config.ESConfig) string {
	//ADDR FORMAT : http://127.0.0.1:3306
	return fmt.Sprintf("%s://%s:%s", cfg.Protocol, cfg.Host, cfg.Port)
}
