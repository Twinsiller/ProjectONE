package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	Id           uint           `json:"id" gorm:"primaryKey"`
	Nickname     string         `json:"nickname" gorm:"type:varchar(30);not null;unique"`
	HashPassword string         `json:"hashpassword" gorm:"type:text;not null"`
	Status       bool           `json:"status" gorm:"default:true"`
	AccessLevel  uint8          `json:"accesslevel" gorm:"default:1;index"`
	Firstname    string         `json:"firstname" gorm:"type:varchar(100);not null"`
	Lastname     string         `json:"lastname" gorm:"type:varchar(100);not null"`
	CreatedAt    time.Time      `json:"createdat" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updatedat" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deletedat" gorm:"index"`
	Posts        []Post         `json:"posts" gorm:"foreignKey:ProfileID"`
	Comments     []Comment      `json:"comments" gorm:"foreignKey:ProfileID"`
}

type ProfileCheck struct {
	Nickname     string `json:"nickname"`
	HashPassword string `json:"hashpassword"`
}
