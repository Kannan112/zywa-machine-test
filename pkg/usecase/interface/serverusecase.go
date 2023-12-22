package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/utils/res"
)

type ServerUseCase interface {
	CardDeliver(ctx context.Context, cardID uint) error
	GetResultCard(ctx context.Context) (res.SampleCardStatus, error)
}
