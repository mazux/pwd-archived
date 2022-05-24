package handler

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/command"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/entity"
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/repository"
)

type SecureLoginHandler struct {
	ProfileRepo repository.ProfileRepo
}

func (h SecureLoginHandler) Handle(cmd command.SecureLogin) error {
	profile, err := h.ProfileRepo.FindByUsername(cmd.ProfileUsername)
	if err != nil {
		return err
	}

	newLogin, err := entity.CreateLogin(cmd.Domain, cmd.Credentials)
	if err != nil {
		return err
	}

	profile.Logins = append(profile.Logins, newLogin)

	return h.ProfileRepo.Save(profile)
}
