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
		body       auth.LoginWithEmailRequest
		setupMocks func(userSvc *MockUserService, tokenSvc, refreshSvc *MockJWTService)
		wantStatus int
		wantBody   any
	}{
		{
			name: " valid login",
			body: auth.LoginWithEmailRequest{Email: "john@example.com", Password: "correct123"},
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

			tt.setupMocks(userSvc, tokenSvc, refreshSvc)

			// Create handler with mocked dependencies
			h := &auth.Handler{
				UserService:         userSvc,
				TokenSevice:         tokenSvc, // note: your original typo is kept
				RefreshTokenService: refreshSvc,
			}

			// Setup router
			server := gin.New()
			group := server.Group("/auth")
			auth.RegisterRoutes(group, h)

			// Make request
			bodyBytes, err := json.Marshal(tt.body)
			assert.Nil(t, err)
			req := httptest.NewRequest(http.MethodPost, "/auth/login-email", strings.NewReader(string(bodyBytes)))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			server.ServeHTTP(w, req)

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

func TestLoginWithPhone(t *testing.T) {
	tests := []struct {
		name       string
		body       auth.LoginWithPhoneRequest
		setupMocks func(userSvc *MockUserService, tokenSvc, refreshSvc *MockJWTService)
		wantStatus int
		wantBody   any
	}{
		{
			name: "valid login",
			body: auth.LoginWithPhoneRequest{Phone: "09188119090", Password: "correct123"},
			setupMocks: func(u *MockUserService, t, r *MockJWTService) {
				user := &entity.User{ID: 42, Phone: utile.StrPtr("09188119090")}
				u.On("LoginWithPhone", "09188119090", "correct123").Return(user, nil)
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

			tt.setupMocks(userSvc, tokenSvc, refreshSvc)

			// Create handler with mocked dependencies
			h := &auth.Handler{
				UserService:         userSvc,
				TokenSevice:         tokenSvc,
				RefreshTokenService: refreshSvc,
			}

			// Setup router
			server := gin.New()
			group := server.Group("/auth")
			auth.RegisterRoutes(group, h)

			// Make request
			bodyBytes, err := json.Marshal(tt.body)
			assert.Nil(t, err)
			req := httptest.NewRequest(http.MethodPost, "/auth/login-phone", strings.NewReader(string(bodyBytes)))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			server.ServeHTTP(w, req)

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
