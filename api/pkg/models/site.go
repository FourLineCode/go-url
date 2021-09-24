package models

import (
	"github.com/lithammer/shortuuid/v3"
)

type Site struct {
	BaseModel

	User   User   `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID uint   `json:"user_id"`
	Key    string `json:"key" gorm:"unique"`
	Url    string `json:"url"`
}

func NewSite(url string, user User) Site {
	key := shortuuid.New()

	return Site{
		Url:  url,
		Key:  key,
		User: user,
	}
}
