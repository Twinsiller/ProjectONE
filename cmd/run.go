package cmd

import (
	v1 "ProjectONE/internal/api/v1"
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/pkg/utils"
)

func Run() error {
	utils.InitLogger("pkg/utils/app.log")

	if err := database.Connect(database.LoadConfigFromEnv()); err != nil {
		return err
	}

	if err := database.Close(); err != nil {
		print("Ошибка при закрытии соединения с базой данных: %v", err)
	}
	print("Соединение с базой данных закрыто")

	router := gin.Default()
	v1.Apies()
	router.Run(":8080")
	return nil
}
