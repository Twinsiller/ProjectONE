package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Получение всех профилей
	router.GET("/profiles", GetProfiles)
	// Получение всех постов
	router.GET("/posts", GetPosts)

	// Получение поста по ID
	router.GET("/profiles/:id", GetProfileByID)
	// Получение профиля по ID
	router.GET("/posts/:id", GetPostByID)

	// Создание нового профиля
	router.POST("/profiles", CreateProfile)
	// Создание новой поста
	router.POST("/posts", CreatePost)

	// Обновление существующего профиля
	router.PUT("/profiles/:id", UpdateProfile)
	// Обновление существующего поста
	router.PUT("/posts/:id", UpdatePost)

	// Удаление профиля
	router.DELETE("/profiles/:id", DeleteProfile)
	// Удаление поста
	router.DELETE("/posts/:id", DeletePost)

	router.Run(":8080")
}
