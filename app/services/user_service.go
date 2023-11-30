package services

import (
	"startup/app/repositories"
	"startup/app/structs"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(request structs.RegisterRequest) (structs.User, error)
	Login(request structs.LoginRequest) (structs.User, error)
	IsEmailAvailable(request structs.CheckEmailRequest) (bool, error)
	SaveAvatar(request structs.UploadAvatarRequest) (structs.User, error)
	GetUserByID(id int) (structs.User, error)
}

type userService struct {
	userRepo    repositories.UserRepository
	roleService RoleService
}

func NewUserService(userRepo repositories.UserRepository, roleService RoleService) *userService {
	return &userService{userRepo, roleService}
}

func (s *userService) Register(request structs.RegisterRequest) (structs.User, error) {
	user := structs.User{}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	role, err := s.roleService.GetRoleByName("user")
	if err != nil {
		return structs.User{}, err
	}

	user.Name = request.Name
	user.Occupation = request.Occupation
	user.Email = request.Email
	user.PasswordHash = string(passwordHash)
	user.RoleID = role.ID

	newUser, err := s.userRepo.Save(user)
	if err != nil {
		return newUser, err
	}

	getUser, err := s.userRepo.FindByID(newUser.ID)
	if err != nil {
		return getUser, err
	}

	return getUser, nil
}

func (s *userService) Login(request structs.LoginRequest) (structs.User, error) {
	user, err := s.userRepo.FindByEmail(request.Email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) IsEmailAvailable(request structs.CheckEmailRequest) (bool, error) {
	_, err := s.userRepo.FindByEmail(request.Email)
	if err != nil {
		return true, err
	}

	return false, err
}

func (s *userService) SaveAvatar(request structs.UploadAvatarRequest) (structs.User, error) {
	user, err := s.userRepo.FindByID(request.User.ID)
	if err != nil {
		return user, err
	}

	user.FileID = &request.FileID

	userUpdated, err := s.userRepo.Update(user)
	if err != nil {
		return userUpdated, err
	}

	return userUpdated, nil
}

func (s *userService) GetUserByID(id int) (structs.User, error) {
	user, err := s.userRepo.FindByID(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
