package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type FileRepository interface {
	Create(file structs.File) (structs.File, error)
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{db}
}

func (r *fileRepository) Create(file structs.File) (structs.File, error) {
	err := r.db.Create(&file).Error

	if err != nil {
		return file, err
	}

	return file, nil
}