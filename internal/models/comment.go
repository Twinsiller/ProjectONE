package models

import "time"

type Comment struct {
	Id               uint      `json:"id" gorm:"primaryKey"`
	Text             string    `json:"text_comment" gorm:"type:text;not null"`
	ProfileID        uint      `json:"profile_id"` // Внешний ключ для профиля
	PostID           uint      `json:"post_id"`    // Внешний ключ для поста
	DatePublication  time.Time `json:"date_publication" gorm:"autoCreateTime"`
	DateLastModified time.Time `json:"date_last_modified" gorm:"autoUpdateTime"`
	Profile          *Profile  `json:"profile" gorm:"foreignKey:ProfileID"`
	Post             *Post     `json:"post" gorm:"foreignKey:PostID"`
}
