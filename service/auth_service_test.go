package service_test

import (
	"testing"

	"authentication/mocks"
	"authentication/model"
	"authentication/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var u = model.User{
	Username: "erwindo",
	Password: "password",
	Name:     "Erwindo Sianipar",
}

func TestCheckUsername(t *testing.T) {
	t.Run("test normal case service check username", func(t *testing.T) {
		authRepoMock := new(mocks.AuthRepoMock)
		authRepoMock.On("CheckUsername", mock.AnythingOfType("string")).Return(nil)

		authService := service.NewAuthService(authRepoMock)
		available := authService.CheckUsername(u.Username)

		t.Run("test username is available", func(t *testing.T) {
			assert.Equal(t, true, available)
		})
	})
}

func TestRegister(t *testing.T) {
	t.Run("test normal case service register", func(t *testing.T) {
		authRepoMock := new(mocks.AuthRepoMock)
		authRepoMock.On("Register", mock.AnythingOfType("*model.User")).Return(nil)

		authService := service.NewAuthService(authRepoMock)
		err := authService.Register(&u)

		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
}

func TestLogin(t *testing.T) {
	password := "$2a$10$fk9IPSmo/VYhu5VJm.vPy.5.XVowBHU3otSDAzTBpMR3YpX2cqYwW"

	t.Run("test normal case service login", func(t *testing.T) {
		authRepoMock := new(mocks.AuthRepoMock)
		authRepoMock.On("Login", mock.AnythingOfType("string")).Return(nil)

		authService := service.NewAuthService(authRepoMock)
		hashedPassword, err := authService.Login(u.Username)
		assert.NoError(t, err)

		t.Run("test get stored password by username is hashed", func(t *testing.T) {
			assert.Equal(t, password, hashedPassword)
		})
	})
}

func TestCheckID(t *testing.T) {
	t.Run("test normal case service check id", func(t *testing.T) {
		authRepoMock := new(mocks.AuthRepoMock)
		authRepoMock.On("CheckID", mock.AnythingOfType("int")).Return(nil)

		authService := service.NewAuthService(authRepoMock)
		available := authService.CheckID(1)

		t.Run("test id is exist for case delete", func(t *testing.T) {
			assert.Equal(t, true, available)
		})
	})
}

func TestDelete(t *testing.T) {
	t.Run("test normal case service delete", func(t *testing.T) {
		authRepoMock := new(mocks.AuthRepoMock)
		authRepoMock.On("Delete", mock.AnythingOfType("int")).Return(nil)

		authService := service.NewAuthService(authRepoMock)
		err := authService.Delete(1)

		t.Run("test data deleted with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
}
