package services

import (
	"EniQilo/entities"
	"EniQilo/repositories"
	"errors"
	"fmt"
)

type UserService interface {
	Create(signupRequest entities.UserRequest) (entities.User, error)
	FindById(id int) (entities.User, error)
	FindAll(filterParams map[string]interface{}) ([]entities.UserResponse, error)
	FindByPhone(phone string) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Create(userRequest entities.UserRequest) (entities.User, error) {
	user := entities.User{
		Name:  userRequest.Name,
		Phone: userRequest.Phone,
		Role:  false,
	}

	newUser, err := s.userRepository.Create(user)
	return newUser, err

}

func (s *userService) FindById(id int) (entities.User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return entities.User{}, err
	}

	// Jika kucing tidak ditemukan, kembalikan error
	if user.Id == 0 {
		return entities.User{}, errors.New("User not found")
	}

	return user, nil
}

func (s *userService) FindByPhone(phone string) (entities.User, error) {
	user, err := s.userRepository.FindByPhone(phone)
	if err != nil {
		return entities.User{}, err
	}

	// Jika kucing tidak ditemukan, kembalikan error
	if user.Id == 0 {
		return entities.User{}, errors.New("User not found")
	}

	return user, nil
}

func (s *userService) FindAll(filterParams map[string]interface{}) ([]entities.UserResponse, error) {
	fmt.Println(filterParams)
	users, err := s.userRepository.FindAll(filterParams)
	if err != nil {
		return nil, err
	}
	return users, nil
}
