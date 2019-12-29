package kloudless

import (
	"context"
	"github.com/cymonevo/cloud-api/sdk"
)

type Client interface {
	GetAccounts(ctx context.Context) (interface{}, error)
	//CreateFolder()
	//RetrieveFolderMetadata()
	//UploadFile()
	//RetrieveFileMetadata()
	//DownloadFile()
}

type clientImpl struct {
	client sdk.Client
	APIKey string
}

func New(cfg Config) *clientImpl {
	return &clientImpl{
		client: sdk.New(sdk.Config{
			Timeout: cfg.Timeout,
			URL:     cfg.URL,
		}),
		APIKey: cfg.APIKey,
	}
}
