package service

import (
	"github.com/hasemeneh/poc-login/config"
	"github.com/hasemeneh/poc-login/internal/definition"
	"github.com/hasemeneh/poc-login/internal/repository"
	"github.com/hasemeneh/poc-login/pkg/database"
	"github.com/hasemeneh/poc-login/pkg/redis"
)

type Service struct {
	Cfg     *config.MainConfig
	Usecase *Usecase
}
type Usecase struct {
	Authentication definition.Authentication
	Authorization  definition.Authorization
	Intools        definition.Intools
}
type Interface struct {
}
type Domain struct {
	User       repository.User
	Permission repository.Permission
	Role       repository.Role
}
type Basic struct {
	DB    *database.Store
	Cache redis.RedisInterface
}
