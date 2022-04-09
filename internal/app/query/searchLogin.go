package query

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/value"
)

type SearchLogin struct {
	ProfileUsername string
	Domain          value.NullableString
	Username        value.NullableString
}
