package models

import "time"

type Reaction struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	RatingID  uint    `json:"-"`
	Rating    *Rating `json:"rating,omitempty" gorm:"foreignKey:RatingID;references:ID"`
	OwnerID   uint    `json:"-"`
	Owner     *User   `json:"owner,omitempty" gorm:"foreignKey:OwnerID;references:ID"`
	Type      string  `json:"type"`
}
