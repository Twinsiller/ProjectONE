package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	Id               uint           `json:"id" gorm:"primaryKey"`
	Title            string         `json:"title" gorm:"type:varchar(255);not null"`
	Description      string         `json:"description" gorm:"type:text;not null"`
	Likes            int            `json:"likes" gorm:"default:0"`
	ProfileID        uint           `json:"profile_id"` // Внешний ключ для профиля
	DatePublication  time.Time      `json:"date_publication" gorm:"autoCreateTime"`
	DateLastModified time.Time      `json:"date_last_modified" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `json:"deletedat" gorm:"index"`
	Profile          *Profile       `json:"profile" gorm:"foreignKey:ProfileID"`
	Comments         []Comment      `json:"comments" gorm:"foreignKey:PostID"`
}
