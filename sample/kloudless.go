package sample

import (
	"context"

	"github.com/cymonevo/cloud-api/internal/log"
	"github.com/cymonevo/cloud-api/provider"
	"github.com/cymonevo/cloud-api/sdk/kloudless"
)

func KloudlessSample() {
	client := provider.GetKloudlessClient()
	folder, err := client.CreateFolder(context.Background(), kloudless.CreateFolderRequest{
		DestID: "root",
		Name:   "sandbox",
	})
	if err != nil {
		return
	}
	log.Infof("CreateFolder", "success create folder %+v", folder)
	contents, err := client.GetFolderContents(context.Background(), kloudless.GetFolderContentsRequest{
		FolderID: folder.ID,
	})
	if err != nil {
		return
	}
	log.Infof("CreateFolder", "success get folder contents %+v", contents)
}
