package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/utils/res"

	"gorm.io/gorm"
)

type cardDatabase struct {
	DB *gorm.DB
}

func NewCardRepository(DB *gorm.DB) interfaces.CardRepository {
	return &cardDatabase{DB}
}

func (c *cardDatabase) GenerateCard(ctx context.Context, user req.User, userID uint) (uint, error) {
	rand.Seed(time.Now().UnixNano())
	var CardID uint

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, 5)
	for i := 0; i < 5; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	cardNumber := string(result)
	fmt.Println("testing")
	query := `INSERT INTO cards (number,user_id) VALUES ($1,$2) RETURNING id`
	if err := c.DB.Raw(query, cardNumber, userID).Scan(&CardID).Error; err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return CardID, nil
}

func (c *cardDatabase) GetCardDetails(userID uint) (domain.Card, error) {
	var cardDetails domain.Card
	query := `select * from card where user_id=$1`

	if err := c.DB.Raw(query, userID).Scan(&cardDetails).Error; err != nil {
		return cardDetails, err
	}
	return cardDetails, nil
}

func (c *cardDatabase) AddDelivery(ctx context.Context, cardID uint, comments []string) error {

	var cardDeliverDetails []domain.DeliveryDetails

	checkCartCount := `select * from delivery_details where card_id=$1`
	if err := c.DB.Raw(checkCartCount, cardID).Scan(&cardDeliverDetails).Error; err != nil {
		return err
	}
	length := len(cardDeliverDetails)

	if length >= 2 {
		return nil
	}

	// Generate a random comment
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(comments))
	randomComment := comments[randomIndex]

	newDeliveryDetail := domain.DeliveryDetails{
		CardID:           cardID,
		DeliveryAttempts: length + 1,    // Increment the attempts
		DeliveryDate:     time.Now(),    // Set the delivery date to the current time
		Comment:          randomComment, // Set your comment here
	}

	query := `INSERT INTO delivery_details (card_id, delivery_attempts, delivery_date, comment) VALUES ($1, $2, $3, $4)`
	if err := c.DB.Exec(query, newDeliveryDetail.CardID, newDeliveryDetail.DeliveryAttempts, newDeliveryDetail.DeliveryDate, newDeliveryDetail.Comment).Error; err != nil {
		return err
	}
	return c.AddDelivery(ctx, cardID, comments)

}

func (c *cardDatabase) ResultData(ctx context.Context) (res.SampleCardStatus, error) {
	var sampleCards res.SampleCardStatus
	query := `SELECT u.id,number as card_id,u.contact as user_contact,d.delivery_date as time_stamp,comment
	FROM users AS u
	JOIN cards AS c ON u.id = c.user_id
	JOIN delivery_details AS d ON c.ID = d.card_id where d.delivery_attempts=2 or d.comment='DELIVERED';		
	`
	if err := c.DB.Raw(query).Scan(&sampleCards).Error; err != nil {
		return sampleCards, err
	}
	return sampleCards, nil
}
