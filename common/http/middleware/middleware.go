package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"authentication/common/http/response"
	"authentication/common/util/token"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	CORS() gin.HandlerFunc
	AUTH() gin.HandlerFunc
}

type middleware struct {
	secretKey string
}

func NewMiddleware(secretKey string) Middleware {
	return &middleware{secretKey: secretKey}
}

func (m *middleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func (m *middleware) AUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		authBearer := strings.Split(authHeader, " ")

		if len(authBearer) == 2 {
			token, err := token.NewToken(m.secretKey).ValidateToken(authBearer[1])
			if token.Valid {
				c.Next()
			} else {
				response.New(c).Error(http.StatusUnauthorized, fmt.Errorf("invalid authorization token %v", err))
				c.Abort()
			}
		} else {
			response.New(c).Error(http.StatusUnauthorized, errors.New("invalid authorization token"))
			c.Abort()
		}
	}
}
