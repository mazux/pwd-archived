package handler

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/query"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/entity"
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/repository"
)

type SearchLoginHandler struct {
	loginRepo repository.LoginRepo
}

func (h SearchLoginHandler) Handle(cmd query.SearchLogin) ([]*entity.Login, error) {
	return h.loginRepo.FindAllBy(cmd.ProfileUsername, cmd.Domain, cmd.Username)
}
