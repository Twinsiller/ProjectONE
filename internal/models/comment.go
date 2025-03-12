package models

import "time"

type Comment struct {
	Id               int       `json:"id" gorm:"primaryKey"`                             // Первичный ключ
	IdAuthor         int       `json:"id_author" gorm:"not null;index"`                  // Внешний ключ на автора
	IdPost           int       `json:"id_post" gorm:"not null;index"`                    // Внешний ключ на пост
	Text             string    `json:"text_comment" gorm:"type:text;not null"`           // Текст комментария
	DatePublication  time.Time `json:"date_publication" gorm:"autoCreateTime"`           // Дата публикации
	DateLastModified time.Time `json:"date_last_modified" gorm:"autoUpdateTime"`         // Дата последнего изменения
	ParentCommentID  *int      `json:"parent_comment_id"`                                // ID родительского комментария для вложенности (если есть)
	ParentComment    *Comment  `json:"parent_comment" gorm:"foreignKey:ParentCommentID"` // Связь с родительским комментарием
	Replies          []Comment `json:"replies" gorm:"foreignKey:ParentCommentID"`        // Связь с ответами (вложенные комментарии)
	Author           Profile   `json:"author" gorm:"foreignKey:IdAuthor"`                // Связь с автором комментария
	Post             Post      `json:"post" gorm:"foreignKey:IdPost"`                    // Связь с постом
}
