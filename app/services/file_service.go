package services

import (
	"startup/app/repositories"
	"startup/app/structs"
)

type FileService interface {
	SaveFile(request structs.FileStoreRequest) (structs.File, error)
}

type fileService struct {
	fileRepo repositories.FileRepository
}

func NewFileService(fileRepo repositories.FileRepository) *fileService {
	return &fileService{fileRepo}
}

func (s *fileService) SaveFile(request structs.FileStoreRequest) (structs.File, error) {
	file := structs.File{
		Name: request.Name,
		Location: request.Location,
	}

	newFile, err := s.fileRepo.Create(file)
	if err != nil {
		return newFile, err
	}

	return newFile, nil
}