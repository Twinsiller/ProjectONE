package v1

import (
	"ProjectONE/internal/service"

	_ "ProjectONE/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Swagger JSON files
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger UI
)

func Apies() {
	router := gin.Default()
	// Маршрут для Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Создание нового профиля
	router.POST("/register", service.CreateProfile)
	// Проверка профиля
	router.POST("/login", login)

	routerv1 := router.Group("/v1")
	// Получение всех профилей
	//router.GET("/profiles", service.GetProfiles)
	profiles := routerv1.Group("/profiles")
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

	posts := routerv1.Group("/posts")
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

	comments := routerv1.Group("/comments")
	comments.Use(authMiddleware())
	{
		// Получение всех постов
		comments.GET("", service.GetComments)

		// Получение профиля по ID
		comments.GET("/:id", service.GetCommentById)

		// Создание новой поста
		comments.POST("", service.CreateComment)

		// Обновление существующего поста
		comments.PUT("/:id", service.UpdateComment)

		// Удаление поста
		comments.DELETE("/:id", service.DeleteComment)
	}
	router.Run(":8080")
}
