package service

import (
	"realtimeChat/internal/domain"
	"realtimeChat/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	CheckUser(username, password string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) Service {
	return Service{Authorization: NewAuthorizationService(repository.Authorization)}
}
