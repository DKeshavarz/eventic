package auth

type LoginWithPhoneRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginWithEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
type ErrorResponse struct {
	Error string `json:"error"`
	Meta  string `json:"meta,omitempty"`
}