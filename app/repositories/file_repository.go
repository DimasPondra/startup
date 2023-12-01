package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type FileRepository interface {
	FindFileByID(ID int) (structs.File, error)
	Create(file structs.File) (structs.File, error)
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{db}
}

func (r *fileRepository) FindFileByID(ID int) (structs.File, error) {
	var file structs.File

	err := r.db.First(&file, ID).Error

	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *fileRepository) Create(file structs.File) (structs.File, error) {
	err := r.db.Create(&file).Error

	if err != nil {
		return file, err
	}

	return file, nil
}
