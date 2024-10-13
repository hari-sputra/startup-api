package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
	GetUserById(id int) (User, error)
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "user"

	createUser, err := s.userRepository.Save(user)

	if err != nil {
		return createUser, err
	}

	return createUser, nil
}

func (s *userService) LoginUser(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User with email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *userService) SaveAvatar(id int, fileLocation string) (User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	updateUser, err := s.userRepository.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (s *userService) GetUserById(id int) (User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User with email not found")
	}

	return user, nil
}
