package storage

import "github.com/jaredmyers/gogym/api/testing/structure/models"

type Storage interface {
	Get(int) *models.User
}
