package user

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	IDColumnName       = "id"
	EmailColumnName    = "email"
	PasswordColumnName = "password"
)

const (
	IDColumnIdx       = 0
	EmailColumnIdx    = 1
	PasswordColumnIdx = 2
)
