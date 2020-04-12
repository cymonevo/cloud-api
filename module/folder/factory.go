package folder

import (
	"github.com/cymonevo/cloud-api/entity"
	"github.com/cymonevo/cloud-api/sdk/kloudless"
)

type Factory interface {
	NewCreateFolderModel(req entity.CreateFolderRequest) *CreateFolderModel
}

type folderFactory struct {
	sdk kloudless.Client
}

func NewFolderFactory(sdk kloudless.Client) Factory {
	return &folderFactory{
		sdk: sdk,
	}
}

func (f *folderFactory) NewCreateFolderModel(req entity.CreateFolderRequest) *CreateFolderModel {
	return &CreateFolderModel{
		sdk: f.sdk,
		req: req,
	}
}
