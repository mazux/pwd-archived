package handler

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/command"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/repository"
)

type ReplaceLoginHandler struct {
	ProfileRepo repository.ProfileRepo
}

func (h ReplaceLoginHandler) Handle(cmd command.ReplaceLogin) error {
	profile, err := h.ProfileRepo.FindByUsername(cmd.ProfileUsername)
	if err != nil {
		return err
	}

	profile.ReplaceLogin(cmd.Domain, cmd.NewCredentials)

	return h.ProfileRepo.Save(profile)
}
