package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"authentication/common/http/request"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlers(t *testing.T) {
	t.Run("test normal case handlers", func(t *testing.T) {
		tests := []struct {
			name   string
			method func(g *gin.Engine)
			body   string
			code   int
		}{
			{
				name: "test status code and index handler",
				method: func(g *gin.Engine) {
					g.GET("/", request.DefaultHandler().Index)
				},
				body: `{"code":200,"message":"application running"}`,
				code: http.StatusOK,
			},
			{
				name: "test status code and no route handler",
				method: func(g *gin.Engine) {
					g.NoRoute(request.DefaultHandler().NoRoute)
				},
				body: `{"code":404,"message":"route not found"}`,
				code: http.StatusNotFound,
			},
		}

		for _, unit := range tests {
			gin := gin.New()
			rec := httptest.NewRecorder()
			unit.method(gin)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			gin.ServeHTTP(rec, req)

			t.Run(unit.name, func(t *testing.T) {
				assert.Equal(t, unit.code, rec.Code)
				assert.Equal(t, unit.body, rec.Body.String())
			})
		}
	})
}
