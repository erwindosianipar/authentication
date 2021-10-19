package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"authentication/common/http/middleware"
	"authentication/common/http/request"
	"authentication/common/util/token"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	secretKey = "v3ry53cr3tk3yk3y!"
	authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImVyd2luZG8iLCJyb2xlIjoiVVNFUiJ9.w8CMUuiIf6T1CT8qcwMgI4oLmzhcr2p2tbEbOg-yecM"
)

func TestCORS(t *testing.T) {
	t.Run("test normal case cors", func(t *testing.T) {
		gin := gin.New()
		rec := httptest.NewRecorder()
		h := request.DefaultHandler()

		gin.Use(middleware.NewMiddleware(secretKey).CORS())
		gin.GET("/", h.Index)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		gin.ServeHTTP(rec, req)

		t.Run("test status code and access allow origin", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
		})
	})
}

func TestAUTH(t *testing.T) {
	t.Run("test normal case auth", func(t *testing.T) {
		gin := gin.New()
		rec := httptest.NewRecorder()
		h := request.DefaultHandler()

		gin.Use(middleware.NewMiddleware(secretKey).AUTH())
		gin.GET("/", h.Index)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

		gin.ServeHTTP(rec, req)

		authHeader := req.Header.Get("Authorization")
		authBearer := strings.Fields(authHeader)

		t.Run("test status code and validate authorization token", func(t *testing.T) {
			if len(authBearer) == 2 {
				token, err := token.NewToken(secretKey).ValidateToken(authBearer[1])
				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(t, true, token.Valid)
				assert.Equal(t, http.StatusOK, rec.Code)
			} else {
				t.Fatal("test failed")
			}
		})
	})
}
