package auth

import (
	"api-golearn/v1/entity"

	"gorm.io/gorm"
)

type Auth interface {
	Login(username string) ([]entity.User, error)
}

type AuthRepo struct {
	db *gorm.DB
}

func NewRepoAuth(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (r *AuthRepo) Login(username string) ([]entity.User, error) {
	var responseLogin []entity.User

	err := r.db.Select("id, username, password, email, created_at").Where("username = ? ", username).Find(&responseLogin).Error
	return responseLogin, err
}
