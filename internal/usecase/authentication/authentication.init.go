package authentication

import (
	"github.com/hasemeneh/poc-login/config"
	"github.com/hasemeneh/poc-login/internal/definition"
	"github.com/hasemeneh/poc-login/internal/repository"
)

type module struct {
	Cfg        *config.MainConfig
	UserRepo   repository.User
	permission repository.Permission
	role       repository.Role
}
type Args struct {
	Cfg        *config.MainConfig
	UserRepo   repository.User
	Permission repository.Permission
	Role       repository.Role
}

func New(o *Args) definition.Authentication {
	return &module{
		UserRepo:   o.UserRepo,
		Cfg:        o.Cfg,
		permission: o.Permission,
		role:       o.Role,
	}
}
