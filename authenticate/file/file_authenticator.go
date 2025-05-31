package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"rbac-project/authenticate"
	"rbac-project/user"
	"strconv"
)

const (
	dbPath    = "./database/file/"
	userTable = "users.csv"
)

type fileAuthenticator struct {
	secretKey string
	source    string
}

func NewFileAuthenticator() authenticate.Authenticator {
	return &fileAuthenticator{
		secretKey: "secret-key",
		source:    "rbac-project",
	}
}

func (a *fileAuthenticator) Login(req authenticate.LoginRequest) (string, error) {
	user, err := a.searchByEmailAndPassword(req.Email, req.Password)
	if err != nil {
		fmt.Println("user not found")
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"source":  a.source,
		"user_id": user.ID,
	})

	return jwtToken.SignedString([]byte(a.secretKey))
}

func (a *fileAuthenticator) searchByEmailAndPassword(email, password string) (*user.User, error) {
	records, err := a.read()
	if err != nil {
		return nil, err
	}

	for i, row := range records {
		if i == 0 {
			continue
		}

		if row[user.EmailColumnIdx] == email && row[user.PasswordColumnIdx] == password {
			id, err := strconv.ParseInt(row[user.IDColumnIdx], 10, 64)
			if err != nil {
				return nil, err
			}

			return &user.User{
				ID:       id,
				Email:    row[user.EmailColumnIdx],
				Password: row[user.PasswordColumnIdx],
			}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("user not found by email and password %v", email))
}

func (a *fileAuthenticator) Authenticate(token string) (int64, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if !jwtToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	if claims["source"] != a.source {
		return 0, errors.New("invalid token")
	}

	return int64(claims["user_id"].(float64)), nil
}

func (a *fileAuthenticator) searchByID(id int64) (*user.User, error) {
	records, err := a.read()
	if err != nil {
		return nil, err
	}

	for i, row := range records {
		if i == 0 {
			continue
		}

		if row[user.IDColumnIdx] == strconv.FormatInt(id, 10) {
			return &user.User{
				ID:       id,
				Email:    row[user.EmailColumnIdx],
				Password: row[user.PasswordColumnIdx],
			}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("user not found by id %v", id))
}

func (a *fileAuthenticator) read() ([][]string, error) {
	f, err := os.Open(dbPath + userTable)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open file error: %v", err))
	}

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("csv read error: %v", err))
	}

	return records, nil
}
