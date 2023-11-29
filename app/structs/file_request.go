package structs

import (
	"mime/multipart"
)

type FileUploadRequest struct {
	Directory string `form:"directory" validate:"required,lowercase"`
	Files []*multipart.FileHeader `form:"files[]" validate:"required,image_type"`
}