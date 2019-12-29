package kloudless

import (
	"context"
	"fmt"

	"github.com/cymonevo/cloud-api/internal/util"
	"github.com/cymonevo/cloud-api/sdk"
)

type Client interface {
	GetAccounts(ctx context.Context) (interface{}, error)
	CreateFolder(ctx context.Context, req CreateFolderRequest) (CreateFolderResponse, error)
	GetFolderContents(ctx context.Context, req GetFolderContentsRequest) (GetFolderContentsResponse, error)
	//RetrieveFolderMetadata()
	//UploadFile()
	//RetrieveFileMetadata()
	//DownloadFile()
}

type Config struct {
	sdk.Config
	APIKey string
	AccID  int
}

type clientImpl struct {
	client sdk.Client
	APIKey string
	AccID  int
}

func New(cfg Config) *clientImpl {
	return &clientImpl{
		client: sdk.New(sdk.Config{
			Timeout: cfg.Timeout,
			URL:     cfg.URL,
		}),
		APIKey: cfg.APIKey,
		AccID:  cfg.AccID,
	}
}

func (c *clientImpl) headers(headers map[string]string) map[string]string {
	return util.CombineMapString(map[string]string{
		"Authorization": fmt.Sprint("APIKey ", c.APIKey),
	}, headers)
}
