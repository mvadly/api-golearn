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

	if pagination.Limit == 0 {
		pagination.Limit = 5
	}

	if pagination.Sort == "" {
		pagination.Sort = "created_at DESC"
	}

	offset := (pagination.Page - 1) * pagination.Limit
	psearch := "%" + pagination.Search + "%"
	err := r.db.Debug().
		Limit(pagination.Limit).
		Offset(offset).
		Where("username LIKE ? OR email LIKE ? OR name LIKE ? ", psearch, psearch, psearch).
		Order(pagination.Sort).
		Find(&users).Error
	return users, err
}

func (r *UserRepo) InsertUser(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
