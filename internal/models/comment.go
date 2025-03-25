package models

import "time"

type Comment struct {
	Id               int       `json:"id" gorm:"primaryKey"`
	IdAuthor         int       `json:"id_author" gorm:"not null;index"`
	IdPost           int       `json:"id_post" gorm:"not null;index"`
	Text             string    `json:"text_comment" gorm:"type:text;not null"`
	DatePublication  time.Time `json:"date_publication" gorm:"autoCreateTime"`
	DateLastModified time.Time `json:"date_last_modified" gorm:"autoUpdateTime"`

	// Вложенность (изменено)
	// ParentCommentID *int      `json:"parent_comment_id"`
	// Replies         []Comment `json:"-" gorm:"foreignKey:ParentCommentID"` // ОТКЛЮЧАЕМ JSON В `Replies`

	// Связи с профилем и постом (исправлено)
	Author Profile `json:"author" gorm:"foreignKey:IdAuthor;constraint:OnDelete:CASCADE;"`
	Post   Post    `json:"post" gorm:"foreignKey:IdPost;constraint:OnDelete:CASCADE;"`
}
