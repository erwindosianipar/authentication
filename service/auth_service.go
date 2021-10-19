package service

import (
	"authentication/model"
	"authentication/repo"
)

type AuthService interface {
	CheckUsername(username string) bool
	Register(user *model.User) error
	Login(username string) (string, error)
	CheckID(id int) bool
	Delete(id int) error
}

type authService struct {
	authRepo repo.AuthRepo
}

func NewAuthService(authRepo repo.AuthRepo) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) CheckUsername(username string) bool {
	return s.authRepo.CheckUsername(username)
}

func (s *authService) Register(user *model.User) error {
	if err := s.authRepo.Register(user); err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(username string) (string, error) {
	password, err := s.authRepo.Login(username)
	if err != nil {
		return "", err
	}

	return password, nil
}

func (s *authService) CheckID(id int) bool {
	return s.authRepo.CheckID(id)
}

func (s *authService) Delete(id int) error {
	if err := s.authRepo.Delete(id); err != nil {
		return err
	}

	return nil
}
