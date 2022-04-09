package handler

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/command"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/repository"
)

type ReplaceLoginHandler struct {
	profileRepo repository.ProfileRepo
}

func (h ReplaceLoginHandler) Handle(cmd command.ReplaceLogin) error {
	profile, err := h.profileRepo.FindByUsername(cmd.ProfileUsername)
	if err != nil {
		return err
	}

	profile.ReplaceLogin(cmd.Domain, cmd.NewCredentials)

	return h.profileRepo.Save(profile)
}
