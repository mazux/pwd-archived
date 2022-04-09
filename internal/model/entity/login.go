package entity

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/value"
)

type Login struct {
	Domain   string         `json:"domain"`
	Username string         `json:"username"`
	Password value.Password `json:"-"`
}

func CreateLogin(domain string, credentials value.Credentials) (*Login, error) {
	encryptedPassword, err := value.Encrypt(credentials.Password, value.NewKey(credentials.Username))
	if err != nil {
		return nil, err
	}

	return &Login{domain, credentials.Username, encryptedPassword}, nil
}
