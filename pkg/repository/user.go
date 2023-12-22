package repository

import (
	"context"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/req"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}
func (c *userDatabase) Create(ctx context.Context, user req.User) (uint, error) {
	var id uint
	if err := c.DB.Raw("INSERT INTO users (contact) VALUES (?) RETURNING id", user.Contact).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}
