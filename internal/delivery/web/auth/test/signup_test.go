package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DKeshavarz/eventic/internal/delivery/web/auth"
	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/pkg/utile"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	tests := []struct {
		tag        string
		body       auth.SignupRequest
		setupMocks func(signupTokenService *MockSignupToken)
		wantStatus int
		wantBody   any
	}{
		{
			tag: "Valid signUp",
			body: auth.SignupRequest{
				Username: "ali", Password: "1234",
				Email: utile.StrPtr("Someone@gmail.com"),
				Token: "Bad.token",
			},
			setupMocks: func(s *MockSignupToken) {
				s.On("Validate", "Bad.token").Return(nil, jwt.ErrInvalidToken)
			},
			wantStatus: http.StatusUnauthorized,
			wantBody:   auth.ErrorResponse{Error: "مشکلی پیش آمده", Meta: jwt.ErrInvalidToken.Error()},
		},
		{
			tag: "Mismatch claims",
			body: auth.SignupRequest{
				Username: "ali", Password: "1234",
				Email: utile.StrPtr("Someone@gmail.com"),
				Token: "my.token",
			},
			setupMocks: func(s *MockSignupToken) {
				s.On("Validate", "my.token").Return(&jwt.SignupTokenClaims{
					Email: "Someone@gmail.com",
				}, nil)
			},
			wantStatus: http.StatusUnauthorized,
			wantBody: auth.ErrorResponse{
				Error: "مشکلی پیش آمده",
				Meta:  "token doesn't match the credential",
			},
		},
		{
			tag: "No email",
			body: auth.SignupRequest{
				Username: "ali", Password: "1234",
				Email: nil,
				Token: "otherguy.token",
			},
			setupMocks: func(s *MockSignupToken) {
				s.On("Validate", "otherguy.token").Return(&jwt.SignupTokenClaims{
					Email: "otherguy@gmail.com",
				}, nil)
			},
			wantStatus: http.StatusUnauthorized,
			wantBody: auth.ErrorResponse{
				Error: "مشکلی پیش آمده",
				Meta:  "token doesn't match the credential",
			},
		},
		{
			tag: "Bad user info",
			body: auth.SignupRequest{
				Username: "ali", Password: "1234",
				Email: utile.StrPtr("Someone@gmail.com"),
				Token: "otherguy.token",
			},
			setupMocks: func(s *MockSignupToken) {
				s.On("Validate", "otherguy.token").Return(&jwt.SignupTokenClaims{
					Email: "Someone@gmail.com",
				}, nil)
			},
			wantStatus: http.StatusUnauthorized,
			wantBody: auth.ErrorResponse{
				Error: "مشکلی پیش آمده",
				Meta:  "token doesn't match the credential",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.tag, func(t *testing.T) {

			tc.setupMocks(signuptoken)

			server := gin.New()
			group := server.Group("/auth")
			h := &auth.Handler{
				SignupToken: signuptoken,
			}
			auth.RegisterRoutes(group, h)

			// Make request
			bodyBytes, err := json.Marshal(tc.body)
			assert.Nil(t, err)
			req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(string(bodyBytes)))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			server.ServeHTTP(w, req)

			// Doing tests
			assert.Equal(t, tc.wantStatus, w.Code)

			if tc.wantStatus == http.StatusOK {
				var got auth.SignUpResponse
				json.Unmarshal(w.Body.Bytes(), &got)
				assert.Equal(t, tc.wantBody, got)
			} else {
				var got auth.ErrorResponse
				json.Unmarshal(w.Body.Bytes(), &got)
				assert.Contains(t, got.Error, (tc.wantBody.(auth.ErrorResponse)).Error)
			}

			userSvc.AssertExpectations(t)
			tokenSvc.AssertExpectations(t)
			refreshSvc.AssertExpectations(t)
		})
	}
}
