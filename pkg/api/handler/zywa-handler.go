package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/res"
)

type ZywaHandler struct {
	userUseCase   services.UserUseCase
	serverUseCase services.ServerUseCase
}

func NewUserHandler(usecase services.UserUseCase, server services.ServerUseCase) *ZywaHandler {
	return &ZywaHandler{
		userUseCase:   usecase,
		serverUseCase: server,
	}
}

// CreateUser creates a new user account and delivers a card (maximum 2 deliveries allowed).
// @Summary Create a user account and deliver a card
// @Description Create a new user account and deliver a card with a maximum of 2 delivery attempts.
// @Tags Users
// @Accept json
// @Produce json
// @Param body body req.User true "User object to be created"
// @Success 200 {object} res.Response "User account created && card delivered to max 2"
// @Failure 422 {object} res.Response "Failed to create user account or deliver card"
// @Router /api/user [post]
func (c *ZywaHandler) CreateUser(ctx *gin.Context) {
	var User req.User
	if err := ctx.BindJSON(&User); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	cartID, err := c.userUseCase.CreateAcc(ctx, User)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "failed to create user account",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := c.serverUseCase.CardDeliver(ctx, cartID); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "failed to deliver card",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Your account has been successfully created, AND check you card status",
		Data:       nil,
		Errors:     nil,
	})
	return

}

func (c *ZywaHandler) GetCardDetails(ctx *gin.Context) {
	card, err := c.serverUseCase.GetResultCard(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "failed to get card details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "CardSample",
		Data:       card,
		Errors:     nil,
	})

}
