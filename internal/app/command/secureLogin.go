package command

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/value"
)

type SecureLogin struct {
	ProfileUsername string
	Domain          string
	Credentials     value.Credentials
}
