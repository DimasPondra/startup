package repositories

import (
	"startup/app/structs"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll() ([]structs.Role, error)
	FindByName(name string) (structs.Role, error)
	Create(role structs.Role) (structs.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) FindAll() ([]structs.Role, error) {
	var roles []structs.Role

	err := r.db.Find(&roles).Error

	if err != nil {
		return roles, err
	}

	return roles, nil
}

func (r *roleRepository) FindByName(name string) (structs.Role, error) {
	var role structs.Role

	err := r.db.Where("name = ?", name).First(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *roleRepository) Create(role structs.Role) (structs.Role, error) {
	err := r.db.Create(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}
