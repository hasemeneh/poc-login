package webservice

import "github.com/hasemeneh/poc-login/pkg/router"

type WebService interface {
	Run() error
	GracefulStop() error
}
type WebRegistrator interface {
	Register(router router.Registrator)
}
