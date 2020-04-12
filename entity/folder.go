package entity

type CreateFolderRequest struct {
	DestID string `json:"dest_id"`
	Name   string `json:"name"`
}

type CreateFolderResponse struct {
	FolderID string `json:"id"`
	Response
}
