package services

import (
	"startup/app/repositories"
	"startup/app/structs"
)

type RoleService interface {
	GetRoles() ([]structs.Role, error)
	GetRoleByName(name string) (structs.Role, error)
	CreateRole(request structs.RoleStoreRequest) (structs.Role, error)
}

type roleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) *roleService {
	return &roleService{roleRepo}
}

func (s *roleService) GetRoles() ([]structs.Role, error) {
	roles, err := s.roleRepo.FindAll()
	if err != nil {
		return roles, err
	}

	return roles, nil
}

func (s *roleService) GetRoleByName(name string) (structs.Role, error) {
	role, err := s.roleRepo.FindByName(name)
	if err != nil {
		return role, err
	}

	return role, nil
}

func (s *roleService) CreateRole(request structs.RoleStoreRequest) (structs.Role, error) {
	role := structs.Role{
		Name: request.Name,
	}

	newRole, err := s.roleRepo.Create(role)
	if err != nil {
		return newRole, err
	}

	return newRole, nil
}