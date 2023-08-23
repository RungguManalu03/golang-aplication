package campaign

import (
	"goaplication/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {

	ID 					string `json:"id" gorm:"unique;default:gen_random_uuid()"`
	UserID 				string
	Name 				string
	ShortDescription 	string
	Description 		string
	Perks				string
	BackerCount			int64
	GoalAmount 			int
	CurrentAmount 		int
	Slug				string
	CreatedAt			time.Time
	UpdatedAt			time.Time
	CampaignImages		[]CampaignImages
	User				user.User
}

func (campaign *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
    campaign.ID = uuid.New().String()
    return
}

type CampaignImages struct {
	ID 					string `json:"id" gorm:"unique;default:gen_random_uuid()"`
	CampaignID 			string
	FileName 			string
	IsPrimary			int
	CreatedAt			time.Time
	UpdatedAt			time.Time
}

func (campaignImages *CampaignImages) BeforeCreate(tx *gorm.DB) (err error) {
    campaignImages.ID = uuid.New().String()
    return
}