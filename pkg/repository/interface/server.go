package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/res"
)

type CardRepository interface {
	GenerateCard(ctx context.Context, user req.User, userID uint) (uint, error)
	GetCardDetails(userID uint) (domain.Card, error)
	AddDelivery(ctx context.Context, cardID uint, comments []string) error
	ResultData(ctx context.Context) (res.SampleCardStatus, error)
}
