package transaction

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID 					string `json:"id" gorm:"unique;default:gen_random_uuid()"`
	CampaignID 			string
	UserID 				string
	Amount 				int
	Status 				string
	Code 				string
	CreatedAt			time.Time
	UpdatedAt			time.Time
}

func (campaign *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	campaign.ID = uuid.New().String()
	return
}