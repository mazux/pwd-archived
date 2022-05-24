package handler

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/query"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model/entity"
	"github.com/MAZEN-Kenjrawi/pwd/internal/model/repository"
)

type SearchLoginHandler struct {
	LoginRepo repository.LoginRepo
}

func (h SearchLoginHandler) Handle(cmd query.SearchLogin) ([]*entity.Login, error) {
	return h.LoginRepo.FindAllBy(cmd.ProfileUsername, cmd.Domain, cmd.Username)
}
