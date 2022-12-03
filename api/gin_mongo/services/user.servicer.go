package services

import "github.com/jaredmyers/gogym/api/gin_mongo/models"

type UserServicer interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
