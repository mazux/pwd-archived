package handler

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/command"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/repository"
)

type DeleteLoginHanlder struct {
	profileRepo repository.ProfileRepo
}

func (h DeleteLoginHanlder) Hanle(cmd command.DeleteLogin) error {
	profile, err := h.profileRepo.FindByUsername(cmd.ProfileUsername)
	if err != nil {
		return err
	}

	return profile.DeleteLogin(cmd.Domain)
}
