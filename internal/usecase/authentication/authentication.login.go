package authentication

import (
	"context"
	"fmt"

	"github.com/hasemeneh/poc-login/internal/models"
)

func (m *module) Login(ctx context.Context, authID, password string) (*models.UserToken, error) {
	user, err := m.UserRepo.GetUserByAuthID(ctx, authID)
	if err != nil {
		return nil, fmt.Errorf("[authentication.Login] failed to get User : %w", err)
	}
	err = m.UserRepo.ValidatePassword(ctx, user, password)
	if err != nil {
		return nil, fmt.Errorf("[authentication.Login] failed to validate password : %w", err)
	}
	resp, err := m.UserRepo.RegisterToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("[authentication.Login] failed to register token : %w", err)
	}
	return resp, nil
}
