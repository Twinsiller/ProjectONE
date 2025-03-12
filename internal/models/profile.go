package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	Id           uint           `json:"id" gorm:"primaryKey"`                                   // Является первичным ключом
	Nickname     string         `json:"nickname" gorm:"type:varchar(100);not null;uniqueIndex"` // Индекс для уникальности nickname
	HashPassword string         `json:"hashpassword" gorm:"type:text;not null"`                 // Хранение пароля как текст
	Status       bool           `json:"status" gorm:"default:true"`                             // Значение по умолчанию true
	AccessLevel  int            `json:"accesslevel" gorm:"default:1;index"`                     // Индекс для AccessLevel
	Firstname    string         `json:"firstname" gorm:"type:varchar(100);not null"`
	Lastname     string         `json:"lastname" gorm:"type:varchar(100);not null"`
	CreatedAt    time.Time      `json:"createdat" gorm:"autoCreateTime"`   // Автоматически ставится время создания
	UpdatedAt    time.Time      `json:"updatedat" gorm:"autoUpdateTime"`   // Автоматически обновляется время при изменении
	DeletedAt    gorm.DeletedAt `json:"deletedat" gorm:"index"`            // Для мягкого удаления с индексом
	Posts        []Post         `json:"posts" gorm:"foreignKey:ProfileID"` // Связь с моделью Post по ProfileID
}

type ProfileCheck struct {
	Nickname     string `json:"nickname"`
	HashPassword string `json:"hashpassword"`
}
