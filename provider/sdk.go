package provider

import (
	"github.com/cymonevo/cloud-api/sdk"
	"sync"

	"github.com/cymonevo/cloud-api/sdk/kloudless"
)

var (
	kloudlessClient     kloudless.Client
	syncKloudlessClient sync.Once
)

func GetKloudlessClient() kloudless.Client {
	if kloudlessClient == nil {
		syncKloudlessClient.Do(func() {
			//TODO: Get from config file
			cfg := kloudless.Config{
				Config: sdk.Config{Timeout: 2, URL: "https://api.kloudless.com/v1"},
				APIKey: "CcRdosbHQLuqzp0HefTUjgiBq__VeMX4dPLyJw86oGRbNPxk",
			}
			kloudlessClient = kloudless.New(cfg)
		})
	}
	return kloudlessClient
}
