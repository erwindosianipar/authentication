package mocks

import (
	"authentication/model"

	"github.com/stretchr/testify/mock"
)

type AuthRepoMock struct {
	mock.Mock
}

func (m *AuthRepoMock) CheckUsername(username string) bool {
	if err := m.Called(username).Error(0); err != nil {
		return false
	}

	return true
}

func (m *AuthRepoMock) Register(user *model.User) error {
	if err := m.Called(user).Error(0); err != nil {
		return err
	}

	return nil
}

func (m *AuthRepoMock) Login(username string) (string, error) {
	if err := m.Called(username).Error(0); err != nil {
		return "", err
	}

	return "$2a$10$fk9IPSmo/VYhu5VJm.vPy.5.XVowBHU3otSDAzTBpMR3YpX2cqYwW", nil
}

func (m *AuthRepoMock) CheckID(id int) bool {
	if err := m.Called(id).Error(0); err != nil {
		return false
	}

	return true
}

func (m *AuthRepoMock) Delete(id int) error {
	if err := m.Called(id).Error(0); err != nil {
		return err
	}

	return nil
}
