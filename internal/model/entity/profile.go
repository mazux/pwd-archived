package entity

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/value"
)

type Profile struct {
	Username string      `json:"username"`
	Secret   *value.Hash `json:"-"`
	Logins   []Login     `json:"logins"`
}

func NewProfile(username, secret string) (*Profile, error) {
	hash, err := value.HashFromString(secret)
	if err != nil {
		return nil, err
	}

	return &Profile{username, hash, []Login{}}, nil
}
