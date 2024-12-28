package cmd

import (
	v1 "ProjectONE/internal/api/v1"
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/pkg/utils"

	"github.com/joho/godotenv"
)

func Run() error {
	// Запуск логгера
	utils.InitLogger("pkg/utils/app.log")

	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		utils.Logger.Fatalf("Error loading .env file")
		return err
	}

	if err := database.Connect(database.LoadConfigFromEnv()); err != nil {
		return err
	}
	defer database.Close()
	v1.Apies()

	return nil
}
