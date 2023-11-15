package services

import (
	"startup/app/repositories"
	"startup/app/structs"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(request structs.RegisterRequest) (structs.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepo}
}

func (s *userService) Register(request structs.RegisterRequest) (structs.User, error) {
	user := structs.User{}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	
	user.Name = request.Name
	user.Occupation = request.Occupation
	user.Email = request.Email
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.userRepo.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}