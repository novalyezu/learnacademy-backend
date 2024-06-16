package user

import (
	"github.com/novalyezu/learnacademy-backend/helper"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
	GetById(id string) (User, error)
}

type userServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (s *userServiceImpl) Register(input RegisterInput) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	ID, errUlid := helper.NewID()
	if errUlid != nil {
		return User{}, errUlid
	}

	user := User{
		ID:        ID,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  string(hashedPassword),
	}

	newUser, errSave := s.userRepository.Save(user)
	if errSave != nil {
		return User{}, errSave
	}

	return newUser, nil
}

func (s *userServiceImpl) Login(input LoginInput) (User, error) {
	user, err := s.userRepository.FindByEmail(input.Email)
	if user.ID == "" {
		return User{}, helper.NewUnauthorizedError("email or password is wrong")
	}
	if err != nil {
		return User{}, err
	}

	errComparePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if errComparePassword != nil {
		return User{}, helper.NewUnauthorizedError("email or password is wrong")
	}

	return user, nil
}

func (s *userServiceImpl) GetById(id string) (User, error) {
	rsUser, err := s.userRepository.FindById(id)
	if err != nil {
		return User{}, err
	}
	return rsUser, nil
}
