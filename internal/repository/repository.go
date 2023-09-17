package repository

import (
	"gorm.io/gorm"
	"realtimeChat/internal/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username string) (int, string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Authorization: NewAuthorizationRepository(db)}
}
