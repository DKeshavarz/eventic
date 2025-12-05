// auth/handler_test.go
package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DKeshavarz/eventic/internal/delivery/web/auth"
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/pkg/utile"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler_LoginWithEmail(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		setupMocks func(userSvc *MockUserService, tokenSvc, refreshSvc *MockJWTService)
		wantStatus int
		wantBody   any // LoginResponse or ErrorResponse
	}{
		{
			name: " valid login",
			body: `{"email":"john@example.com","password":"correct123"}`,
			setupMocks: func(u *MockUserService, t, r *MockJWTService) {
				user := &entity.User{ID: 42, Email: utile.StrPtr("john@example.com")}
				u.On("LoginWithEmail", "john@example.com", "correct123").Return(user, nil)
				t.On("Generate", user).Return("fake.jwt.access.token", nil)
				r.On("Generate", user).Return("fake.jwt.refresh.token", nil)
			},
			wantStatus: http.StatusOK,
			wantBody: auth.LoginResponse{
				Token:        "fake.jwt.access.token",
				RefreshToken: "fake.jwt.refresh.token",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			userSvc := new(MockUserService)
			tokenSvc := new(MockJWTService)
			refreshSvc := new(MockJWTService)

			tt.setupMocks(userSvc, tokenSvc, refreshSvc)

			// Create handler with mocked dependencies
			h := &auth.Handler{
				UserService:         userSvc,
				TokenSevice:         tokenSvc, // note: your original typo is kept
				RefreshTokenService: refreshSvc,
			}

			// Setup router
			r := gin.New()
			group := r.Group("/auth")
			auth.RegisterRoutes(group, h)

			// Make request
			req := httptest.NewRequest(http.MethodPost, "/auth/login-email", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tt.wantStatus, w.Code)

			if tt.wantStatus == http.StatusOK {
				var got auth.LoginResponse
				json.Unmarshal(w.Body.Bytes(), &got)
				assert.Equal(t, tt.wantBody, got)
			} else {
				var got auth.ErrorResponse
				json.Unmarshal(w.Body.Bytes(), &got)
				assert.Contains(t, got.Error, (tt.wantBody.(auth.ErrorResponse)).Error)
			}

			// Verify all expected mock calls happened
			userSvc.AssertExpectations(t)
			tokenSvc.AssertExpectations(t)
			refreshSvc.AssertExpectations(t)
		})
	}
}
