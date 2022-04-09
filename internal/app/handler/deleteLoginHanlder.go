package handler

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/command"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/repository"
)

type DeleteLoginHandler struct {
	profileRepo repository.ProfileRepo
}

func (h DeleteLoginHandler) Handle(cmd command.DeleteLogin) error {
	profile, err := h.profileRepo.FindByUsername(cmd.ProfileUsername)
	if err != nil {
		return err
	}

	return profile.DeleteLogin(cmd.Domain)
}
