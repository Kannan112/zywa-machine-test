package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/utils/req"
)

type UserUseCase interface {
	CreateAcc(ctx context.Context, user req.User) (uint, error)
}
