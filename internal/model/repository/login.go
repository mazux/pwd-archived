package repository

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/entity"
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/value"
)

type LoginRepo interface {
	FindAllBy(string, value.NullableString, value.NullableString) ([]*entity.Login, error)
}
