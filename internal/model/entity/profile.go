package entity

import (
	"fmt"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/value"
)

type Profile struct {
	Username string     `json:"username"`
	Secret   value.Hash `json:"-"`
	Logins   []*Login   `json:"logins"`
}

func (p *Profile) GetLogin(domain string) (*Login, error) {
	for _, login := range p.Logins {
		if login.Domain == domain {
			return login, nil
		}
	}

	return nil, fmt.Errorf("unable to find login information for domain %s", domain)
}

func (p *Profile) ReplaceLogin(domain string, newCredentials value.Credentials) error {
	if err := p.DeleteLogin(domain); err != nil {
		return err
	}

	login, err := CreateLogin(domain, newCredentials)
	if err != nil {
		return err
	}

	p.Logins = append(p.Logins, login)

	return nil
}

func (p *Profile) DeleteLogin(domain string) error {
	for index, login := range p.Logins {
		if login.Domain == domain {
			p.Logins = append([]*Login{}, p.Logins[:index]...)
			p.Logins = append(p.Logins, p.Logins[index+1:]...)

			return nil
		}
	}

	return fmt.Errorf("unable to find login information for domain %s", domain)
}

func NewProfile(username, secret string) (*Profile, error) {
	hash, err := value.HashString(secret)
	if err != nil {
		return nil, err
	}

	return &Profile{username, hash, nil}, nil
}
