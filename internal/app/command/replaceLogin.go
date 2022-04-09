package command

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/value"
)

type ReplaceLogin struct {
	ProfileUsername string
	Domain          string
	NewCredentials  value.Credentials
}
