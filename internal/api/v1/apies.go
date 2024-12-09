package v1

import (
	"ProjectONE/internal/service"

	"github.com/gin-gonic/gin"
)

func Apies() {
	router := gin.Default()

	// Создание нового профиля
	router.POST("/register", service.CreateProfile)
	// Проверка профиля
	router.POST("/login", login)

	// Получение всех профилей
	//router.GET("/profiles", service.GetProfiles)
	profiles := router.Group("/profiles")
	profiles.Use(authMiddleware())
	{
		// Получение всех профилей
		profiles.GET("", service.GetProfiles)

		// Получение поста по ID
		profiles.GET("/:id", service.GetProfileById)

		// Обновление существующего профиля
		profiles.PUT("/:id", service.UpdateProfile)

		// Удаление профиля
		profiles.DELETE("/:id", service.DeleteProfile)
	}

	posts := router.Group("/posts")
	posts.Use(authMiddleware())
	{
		// Получение всех постов
		posts.GET("", service.GetPosts)

		// Получение профиля по ID
		posts.GET("/:id", service.GetPostById)

		// Создание новой поста
		posts.POST("", service.CreatePost)

		// Обновление существующего поста
		posts.PUT("/:id", service.UpdatePost)

		// Удаление поста
		posts.DELETE("/:id", service.DeletePost)
	}
	router.Run(":8080")
}
