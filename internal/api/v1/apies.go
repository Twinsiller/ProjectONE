package v1

import "github.com/gin-gonic/gin"

func Apies() {
	profile := Router.Group("/profiles")
	{
		// Получение всех профилей
		profile.GET("/profiles", service.GetProfiles)

		// // Получение поста по ID
		// profile.GET("/profiles/:id", GetProfileByID)

		// // Создание нового профиля
		// profile.POST("/profiles", CreateProfile)

		// // Обновление существующего профиля
		// profile.PUT("/profiles/:id", UpdateProfile)

		// // Удаление профиля
		// profile.DELETE("/profiles/:id", DeleteProfile)
	}
	post := Router


}
