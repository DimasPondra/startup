package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user structs.User) (structs.User, error)
	FindByEmail(email string) (structs.User, error)
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

func (r *userRepository) FindByEmail(email string) (structs.User, error) {
	var user structs.User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}