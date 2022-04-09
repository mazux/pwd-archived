package handler

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/query"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/entity"
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/repository"
)

type GetLoginHandler struct {
	profileRepo repository.ProfileRepo
}

func (h GetLoginHandler) Handle(cmd query.GetLogin) (*entity.Login, error) {
	profile, err := h.profileRepo.FindByUsername(cmd.ProfileUsername)
	if err != nil {
		return nil, err
	}

	return profile.GetLogin(cmd.Domain)
}
