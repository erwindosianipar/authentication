package v1_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	v1 "authentication/api/v1"
	"authentication/infra"
	"authentication/mocks"
	"authentication/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockUser = model.User{
	Username: "erwindo",
	Password: "password",
	Name:     "Erwindo Sianipar",
}

func TestRegister(t *testing.T) {
	t.Run("test normal case register", func(t *testing.T) {
		authServiceMock := new(mocks.AuthServiceMock)
		authServiceMock.On("CheckUsername", mock.AnythingOfType("string")).Return(nil)
		authServiceMock.On("Register", mock.AnythingOfType("*model.User")).Return(nil)

		gin := gin.New()
		rec := httptest.NewRecorder()

		authHandler := v1.NewAuthHandler(authServiceMock, infra.New("../../config/config.json"))
		gin.POST("/register", authHandler.Register)

		body, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		exp := `{"code":201,"message":"success: user registered"}`

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, exp, rec.Body.String())
		})
	})
}

func TestLogin(t *testing.T) {
	t.Run("test normal case login", func(t *testing.T) {
		authServiceMock := new(mocks.AuthServiceMock)
		authServiceMock.On("Login", mock.AnythingOfType("string")).Return(nil)

		gin := gin.New()
		rec := httptest.NewRecorder()

		authHandler := v1.NewAuthHandler(authServiceMock, infra.New("../../config/config.json"))
		gin.POST("/login", authHandler.Login)

		body, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		var mockResponse model.Token
		err = json.Unmarshal(rec.Body.Bytes(), &mockResponse)
		assert.NoError(t, err)

		exp := string(time.Now().Add(time.Hour * 2).Format(time.RFC3339))

		t.Run("test status code and token expiration", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, exp, mockResponse.Expired)
		})
	})
}

func TestDelete(t *testing.T) {
	t.Run("test normal case delete", func(t *testing.T) {
		authServiceMock := new(mocks.AuthServiceMock)
		authServiceMock.On("CheckID", mock.AnythingOfType("int")).Return(nil)
		authServiceMock.On("Delete", mock.AnythingOfType("int")).Return(nil)

		gin := gin.New()
		rec := httptest.NewRecorder()

		authHandler := v1.NewAuthHandler(authServiceMock, infra.New("../../config/config.json"))
		gin.DELETE("/delete", authHandler.Delete)

		req := httptest.NewRequest(http.MethodDelete, "/delete?id=1", nil)
		gin.ServeHTTP(rec, req)

		exp := `{"code":200,"message":"success: user deleted"}`

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, exp, rec.Body.String())
		})
	})
}
