package authenticate

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
