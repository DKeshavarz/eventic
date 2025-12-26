package middelware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jwtService := jwt.NewTokenService(&jwt.AccessTokenConfig{
		Duration: time.Hour,
		Secret: []byte("moew"),
	})
	otherJwtService := jwt.NewTokenService(&jwt.AccessTokenConfig{
		Duration: time.Hour,
		Secret: []byte("moewww"),
	})
	validToken, err := jwtService.Generate(&entity.User{ID: 5})
	if err != nil {
		t.Fail()
	}
	invalidToken, err := otherJwtService.Generate(&entity.User{ID: 5})
	if err != nil {
		t.Fail()
	}

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Missing Authorization header",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Authorization header required"}`,
		},
		{
			name:           "Invalid Authorization header format",
			authHeader:     "InvalidHeader",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Authorization header format must be Bearer {token}"}`,
		},
		{
			name:           "Invalid token",
			authHeader:     "Bearer invalidtoken",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid or expired token"}`,
		},
		{
			name:           "Valid token",
			authHeader:     fmt.Sprintf("Bearer %s", validToken),
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message": "success"}`,
		},
		{
			name:           "Expired token",
			authHeader:     fmt.Sprintf("Bearer %s", invalidToken),
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid or expired token"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := gin.New()
			r.Use(Auth(jwtService))
			r.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
