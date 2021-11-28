package user

import (
	"api-golearn/v1/entity"

	"gorm.io/gorm"
)

type User interface {
	InsertUser(create entity.User) (entity.User, error)
	GetAllUsers(pagination Pagination) ([]entity.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewRepoUser(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) GetAllUsers(pagination Pagination) ([]entity.User, error) {
	var users []entity.User
	offset := (pagination.Page - 1) * pagination.Limit
	err := r.db.Debug().Limit(pagination.Limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *UserRepo) InsertUser(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
