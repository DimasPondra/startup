package services

import (
	"startup/app/repositories"
	"startup/app/structs"
)

type FileService interface {
	GetFileByID(ID int) (structs.File, error)
	SaveFile(request structs.FileStoreRequest) (structs.File, error)
}

type fileService struct {
	fileRepo repositories.FileRepository
}

func NewFileService(fileRepo repositories.FileRepository) *fileService {
	return &fileService{fileRepo}
}

func (s *fileService) GetFileByID(ID int) (structs.File, error) {
	file, err := s.fileRepo.FindFileByID(ID)
	if err != nil {
		return file, err
	}

	return file, nil
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