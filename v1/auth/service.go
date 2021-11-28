package auth

import (
	"api-golearn/v1/entity"
)

type AuthService interface {
	Login(username string) ([]entity.User, error)
	Profile(id interface{}) ([]entity.User, error)
}

type authService struct {
	authUserService AuthRepo
}

func NewAuthService(authUserService AuthRepo) *authService {
	return &authService{authUserService}
}

func (s *authService) Login(username string) ([]entity.User, error) {
	requestServiceLogin, err := s.authUserService.Login(username)
	return requestServiceLogin, err
}

func (s *authService) Profile(id interface{}) ([]entity.User, error) {
	requestServiceLogin, err := s.authUserService.Profile(id)
	return requestServiceLogin, err
}
