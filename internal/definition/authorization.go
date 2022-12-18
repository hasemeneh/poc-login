package definition

import (
	"context"

	"github.com/hasemeneh/poc-login/internal/models"
)

type Authorization interface {
	GetAllUserAccess(ctx context.Context, token string) ([]models.RolePermission, error)
}
