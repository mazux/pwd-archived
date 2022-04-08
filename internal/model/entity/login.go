package entity

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/value"
)

type Login struct {
	Domain   string         `json:"domain"`
	Username string         `json:"username"`
	Password value.Password `json:"-"`
}

func CreateLogin(domain, username, password string) (*Login, error) {
	encryptedPassword, err := value.Encrypt(password, value.NewKey(username))
	if err != nil {
		return nil, err
	}

	return &Login{domain, username, encryptedPassword}, nil
}
