package campaign

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	ID 					string `json:"id" gorm:"unique;default:gen_random_uuid()"`
	UserID 				uuid.UUID 
	Name 				string
	ShortDescription 	string
	Description 		string
	Perks				string
	BackerCount			int64
	GoalAmount 			int64
	CurrentAmount 		int64
	Slug				string
	CreatedAt			time.Time
	UpdatedAt			time.Time
	CampaignImages		[]CampaignImages
}
type CampaignImages struct {
	ID 					string `json:"id" gorm:"unique;default:gen_random_uuid()"`
	CampaignID 			uuid.UUID 
	FileName 			string
	IsPrimary			int
	CreatedAt			time.Time
	UpdatedAt			time.Time
}

func (u *Campaign) BeforeCreateCampaign(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}