package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID 				string `json:"id" gorm:"unique;default:gen_random_uuid()"`
	Name           	string
	Occupation     	string
	Email          	string
	PasswordHash   	string
	AvatarFileName 	string
	Role           	string
	CreatedAt      	time.Time
	UpdatedAt      	time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New().String()
    return
}