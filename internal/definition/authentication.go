package definition

import (
	"context"

	"github.com/hasemeneh/poc-login/internal/models"
)

type Authentication interface {
	Register(ctx context.Context, req *models.RegisterReq) error
	Login(ctx context.Context, authID, password string) (*models.UserToken, error)
	Logout(ctx context.Context, token string) error
}
