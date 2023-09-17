package service

import (
	"errors"
	"realtimeChat/internal/domain"
	"realtimeChat/internal/repository"
	"realtimeChat/pkg/hashing"
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) CreateUser(user domain.User) (int, error) {
	var err error
	user.Password, err = hashing.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateUser(user)
}

func (s *AuthorizationService) CheckUser(username, password string) (int, error) {
	userID, hash, err := s.repo.GetUser(username)
	if err != nil {
		return 0, err
	}

	ok := hashing.CheckPasswordHash(password, hash)
	if !ok {
		return 0, errors.New("wrong password")
	}

	return userID, err
}
