package repository

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/entity"
)

type ProfileRepo interface {
	FindByUsername(string) (*entity.Profile, error)
	Save(*entity.Profile) error
}
