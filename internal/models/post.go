package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	Id               int            `json:"id" gorm:"type:serial;primaryKey"`                // Является первичным ключом
	Title            string         `json:"title" gorm:"type:varchar(255);not null"`         // Название поста
	IdAuthor         int            `json:"id_author" gorm:"type:integer;index;not null"`    // Индекс для авторов
	Description      string         `json:"description" gorm:"type:text;not null"`           // Описание поста
	Likes            uint           `json:"likes" gorm:"default:0"`                          // Количество лайков
	DatePublication  time.Time      `json:"date_publication" gorm:"autoCreateTime"`          // Дата публикации
	DateLastModified time.Time      `json:"date_last_modified" gorm:"autoUpdateTime"`        // Дата последнего изменения
	DeletedAt        gorm.DeletedAt `json:"deletedat" gorm:"index"`                          // Для мягкого удаления с индексом
	PostProfile      *Profile       `json:"author" gorm:"type:integer;foreignKey:IdProfile"` // Связь с автором
	Comments         []*Comment     `json:"comments" gorm:"foreignKey:PostID"`               // Связь с комментариями, внешний ключ в модели Comment
}
