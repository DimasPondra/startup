package structs

type CampaignImagesUploadRequest struct {
	FileIDs []int `json:"file_ids" validate:"required,min=1,ids_exists_in_files"`
}

type CampaignImageStoreRequest struct {
	IsPrimary  int
	CampaignID int
	FileID     int
}