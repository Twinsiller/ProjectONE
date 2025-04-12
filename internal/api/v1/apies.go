package v1

import (
	"ProjectONE/internal/service"
	"ProjectONE/pkg/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	router.GET("/shutdown")

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

	//router.Run(":8080") // Это прошлое

	// Создаём кастомный сервер
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Запуск сервера в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// Канал для сигналов завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Блокировка до получения сигнала
	<-quit

	if quit ==  

	utils.Logger.Println("Завершение работы сервера...")
	

	// Таймаут для graceful shutdown
	ctx, cancel := context.WithTimeout(
		context.Background(), 
		5*time.Second,
	)
	defer cancel()

	// Остановка сервера
	if err := srv.Shutdown(ctx); err != nil {
		utils.Logger.Fatal("Server forced to shutdown:", err)
	}

	utils.Logger.Println("Server exiting")
}
