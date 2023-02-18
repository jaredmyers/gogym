package storage

import "github.com/jaredmyers/gogym/api/testing/structure/models"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *models.User {
	return &models.User{
		ID:   1,
		Name: "Foo",
	}
}
