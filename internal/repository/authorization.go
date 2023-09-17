package repository

import (
	"gorm.io/gorm"
	"realtimeChat/internal/domain"
)

type AuthorizationRepository struct {
	db *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) CreateUser(user domain.User) (int, error) {
	err := r.db.Create(&user)
	return user.ID, err.Error
}

func (r *AuthorizationRepository) GetUser(username string) (int, string, error) {
	user := domain.User{}
	err := r.db.Where("username = ?", username).First(&user)
	return user.ID, user.Password, err.Error
}
