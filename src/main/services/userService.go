package services

import (
	"github.com/2k4sm/httpCoffee/src/main/dto"
	"github.com/2k4sm/httpCoffee/src/main/entities"
	"github.com/2k4sm/httpCoffee/src/main/repositories"
)

type UserServiceInterface interface {
	GetAll() []entities.User
	GetById(id int) entities.User
	GetByName(name string) entities.User
	GetByEmail(email string) entities.User
	CreateUser(user dto.CreateUser) entities.User
	DeleteUser(id int) (entities.User, error)
}

type userService struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		UserRepository: userRepository,
	}
}

func (u *userService) GetAll() []entities.User {
	return u.UserRepository.FindAll()
}

func (u *userService) GetById(id int) entities.User {
	return u.UserRepository.FindById(uint(id))
}

func (u *userService) GetByName(name string) entities.User {
	return u.UserRepository.FindByName(name)
}

func (u *userService) GetByEmail(email string) entities.User {
	return u.UserRepository.FindByEmail(email)
}

func (u *userService) CreateUser(user dto.CreateUser) entities.User {
	newUser := dto.ParseToUserEntity(user)
	return u.UserRepository.CreateUser(newUser)
}

func (u *userService) DeleteUser(id int) (entities.User, error) {
	return u.UserRepository.DeleteUser(uint(id))
}
