package repositories

import (
	"github.com/2k4sm/httpCoffee/src/main/entities"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAll() []entities.User
	FindById(id uint) entities.User
	FindByName(name string) entities.User
	FindByEmail(email string) entities.User
	CreateUser(user entities.User) entities.User
	DeleteUser(id uint) (entities.User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		Db: db,
	}
}

func (u *userRepository) FindAll() []entities.User {
	var users []entities.User
	u.Db.Find(&users)

	if len(users) == 0 {
		log.Error("No users found")
	}

	log.Infof("Found %d users", len(users))

	return users
}

func (u *userRepository) FindById(id uint) entities.User {
	var user entities.User

	u.Db.First(&user, id)

	if user.ID == 0 {
		log.Infof("No user found with id: %s", id)
	}

	log.Infof("Found user with id: %s", id)

	return user
}

func (u *userRepository) FindByName(name string) entities.User {
	var user entities.User

	u.Db.Where("name = ?", name).First(&user)

	if user.ID == 0 {
		log.Infof("No user found with name: %s", name)
	}

	log.Infof("Found user with name: %s", name)

	return user
}

func (u *userRepository) FindByEmail(email string) entities.User {
	var user entities.User

	u.Db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		log.Infof("No user found with email: %s", email)
	}

	log.Infof("Found user with email: %s", email)

	return user
}

func (u *userRepository) CreateUser(user entities.User) entities.User {
	err := u.Db.Create(&user)

	if err.Error != nil {
		log.Errorf("Error creating user: %s", err.Error)
	}

	log.Infof("Created user with id: %s", user.ID)

	return user
}

func (u *userRepository) DeleteUser(id uint) (entities.User, error) {
	userToDel := entities.User{}

	u.Db.First(&userToDel, id)

	if userToDel.ID == 0 {
		log.Infof("No user found with id: %s", id)
		return userToDel, gorm.ErrRecordNotFound
	}

	u.Db.Delete(&userToDel)

	return userToDel, nil
}
