package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id int) (structs.User, error)
	FindByEmail(email string) (structs.User, error)
	Save(user structs.User) (structs.User, error)
	Update(user structs.User) (structs.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByID(id int) (structs.User, error) {
	var user structs.User

	err := r.db.Preload("Role").First(&user, id).Error

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

func (r *userRepository) Save(user structs.User) (structs.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user structs.User) (structs.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
