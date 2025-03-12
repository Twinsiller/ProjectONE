package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	Id               int            `json:"id" gorm:"primaryKey"`                     // Является первичным ключом
	Title            string         `json:"title" gorm:"type:varchar(255);not null"`  // Название поста
	IdAuthor         int            `json:"id_author" gorm:"index;not null"`          // Индекс для авторов
	Description      string         `json:"description" gorm:"type:text;not null"`    // Описание поста
	Likes            int            `json:"likes" gorm:"default:0"`                   // Количество лайков
	DatePublication  time.Time      `json:"date_publication" gorm:"autoCreateTime"`   // Дата публикации
	DateLastModified time.Time      `json:"date_last_modified" gorm:"autoUpdateTime"` // Дата последнего изменения
	DeletedAt        gorm.DeletedAt `json:"deletedat" gorm:"index"`                   // Для мягкого удаления с индексом
	Author           Profile        `json:"author" gorm:"foreignKey:IdAuthor"`        // Связь с автором
	Comments         []Comment      `json:"comments" gorm:"foreignKey:PostID"`        // Связь с комментариями, внешний ключ в модели Comment
}
