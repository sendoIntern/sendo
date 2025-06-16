package request

type LoginRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}