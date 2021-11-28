package user

import (
	"api-golearn/v1/entity"
	"api-golearn/v1/util"
)

type UserService interface {
	CreateUser(create RequestCreateUser) (entity.User, error)
	GetUsers(pagination Pagination) (entity.User, error)
}

type userService struct {
	userRepoService UserRepo
}

func NewUserService(userRepoService UserRepo) *userService {
	return &userService{userRepoService}
}

func (s *userService) GetUsers(pagination Pagination) ([]entity.User, error) {
	users, err := s.userRepoService.GetAllUsers(pagination)
	return users, err
}

func (s *userService) CreateUser(create RequestCreateUser) (entity.User, error) {
	createUser := entity.User{
		Username: create.Username,
		Password: util.HashAndSalt([]byte(create.Password)),
		Email:    create.Email,
		Name:     create.Name,
	}
	requestServiceCreateUser, err := s.userRepoService.InsertUser(createUser)
	return requestServiceCreateUser, err
}
