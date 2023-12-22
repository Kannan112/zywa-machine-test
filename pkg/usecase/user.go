package usecase

import (
	"context"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/req"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
	cardRepo interfaces.CardRepository
}

func NewUserUseCase(user interfaces.UserRepository, card interfaces.CardRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: user,
		cardRepo: card,
	}
}

func (c *userUseCase) CreateAcc(ctx context.Context, user req.User) (uint, error) {
	id, err := c.userRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}
	cardID, err := c.cardRepo.GenerateCard(ctx, user, uint(id))
	if err != nil {
		return 0, err
	}
	return cardID, nil
}
