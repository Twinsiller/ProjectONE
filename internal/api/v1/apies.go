package v1

import (
	"ProjectONE/internal/service"

	"github.com/gin-gonic/gin"
)

func Apies() {
	router := gin.Default()
	// Получение всех профилей
	//router.GET("/profiles", service.GetProfiles)
	profiles := router.Group("/profiles")
	{
		// Получение всех профилей
		profiles.GET("", service.GetProfiles)

		// Получение поста по ID
		profiles.GET("/:id", service.GetProfileById)

		// Создание нового профиля
		profiles.POST("", service.CreateProfile)

		// Обновление существующего профиля
		profiles.PUT("/:id", service.UpdateProfile)

		// Удаление профиля
		profiles.DELETE("/:id", service.DeleteProfile)
	}
	post := router.Group("/posts")
	{
		// Получение всех постов
		post.GET("", service.GetPosts)

		// Получение профиля по ID
		post.GET("/:id", service.GetPostById)

		// Создание новой поста
		post.POST("", service.CreatePost)

		// Обновление существующего поста
		post.PUT("/:id", service.UpdatePost)

		// Удаление поста
		post.DELETE("/:id", service.DeletePost)
	}
	router.Run(":8080")
}
