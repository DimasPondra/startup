package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user structs.User) (structs.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user structs.User) (structs.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}