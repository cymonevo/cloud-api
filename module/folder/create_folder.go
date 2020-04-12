package folder

import (
	"context"
	"errors"

	"github.com/cymonevo/cloud-api/entity"
	"github.com/cymonevo/cloud-api/internal/log"
	"github.com/cymonevo/cloud-api/sdk/kloudless"
)

const createFolderTag = "Folder|Create"

type CreateFolderModel struct {
	req entity.CreateFolderRequest
	sdk kloudless.Client
}

func (m *CreateFolderModel) Do(ctx context.Context) (entity.CreateFolderResponse, error) {
	var response entity.CreateFolderResponse
	if err := m.Validate(ctx); err != nil {
		log.ErrorDetail(createFolderTag, "error validation: %v", err)
		response.Message = err.Error()
		return response, err
	}
	result, err := m.sdk.CreateFolder(ctx, kloudless.CreateFolderRequest{
		DestID: m.req.DestID,
		Name:   m.req.Name,
	})
	if err != nil {
		log.ErrorDetail(createFolderTag, "error create folder: %v", err)
		response.Message = err.Error()
		return response, err
	}
	response.FolderID = result.ID
	log.Infof(createFolderTag, "success create folder")
	return response, nil
}

func (m *CreateFolderModel) Validate(ctx context.Context) error {
	if m.req.DestID == "" || m.req.Name == "" {
		return errors.New("invalid request")
	}
	return nil
}
