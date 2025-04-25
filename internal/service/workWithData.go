package service

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"encoding/json"
	"os"
	"time"
)

// DumpDataToFile выгружает все записи из таблиц в файл
func DumpDataToFile() error {
	utils.Logger.Info("Начинаем выгрузку данных из БД...")

	var profiles []models.Profile
	var posts []models.Post
	var comments []models.Comment

	if err := database.DbPostgres.Find(&profiles).Error; err != nil {
		return err
	}
	if err := database.DbPostgres.Find(&posts).Error; err != nil {
		return err
	}
	if err := database.DbPostgres.Find(&comments).Error; err != nil {
		return err
	}

	data := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"profiles":  profiles,
		"posts":     posts,
		"comments":  comments,
	}

	file, err := os.Create("dump.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	utils.Logger.Info("Данные успешно выгружены в dump.json")
	return nil
}
