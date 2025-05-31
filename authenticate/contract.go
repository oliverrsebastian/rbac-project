package authenticate

type Authenticator interface {
	Authenticate(token string) (int64, error)
	Login(req LoginRequest) (string, error)
}
