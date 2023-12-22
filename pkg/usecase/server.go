package usecase

import (
	"context"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/res"
)

const (
	Delivered        = "DELIVERED"
	UserNotAvailable = "User not available"
	UserDeclined     = "User declined to accept package"
)

var comments = []string{Delivered, UserNotAvailable, UserDeclined}

type ServerRepo struct {
	CardRepo interfaces.CardRepository
	UserRepo interfaces.UserRepository
}

func NewServerUseCase(cardRepo interfaces.CardRepository, userRepo interfaces.UserRepository) services.ServerUseCase {
	return &ServerRepo{
		CardRepo: cardRepo,
		UserRepo: userRepo,
	}
}

func (c *ServerRepo) CardDeliver(ctx context.Context, cardID uint) error {
	err := c.CardRepo.AddDelivery(ctx, cardID, comments)
	if err != nil {
		return err
	}
	return nil
}

func (c *ServerRepo) GetResultCard(ctx context.Context) (res.SampleCardStatus, error) {
	card, err := c.CardRepo.ResultData(ctx)
	if err != nil {
		return card, err
	}
	return card, nil
}
