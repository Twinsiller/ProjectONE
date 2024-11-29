package database

import (
	"ProjectONE/pkg/utils"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Импортируем PostgreSQL драйвер
)

// DB — глобальная переменная для хранения подключения к базе данных
var DbPostgres *sql.DB

// Config — структура для хранения конфигурации подключения
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

// LoadConfigFromEnv загружает настройки базы данных из переменных окружения
func LoadConfigFromEnv() Config {
	cfg := Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	// utils.Logger.Printf("Проверка загрузки\nuser=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
	// 	cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	return cfg
}

// Connect устанавливает соединение с базой данных
func Connect(cfg Config) error {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)
	utils.Logger.Printf("Проверка подключения\n%s", dsn)
	var err error
	DbPostgres, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Проверяем соединение
	if err = DbPostgres.Ping(); err != nil {
		return fmt.Errorf("база данных недоступна: %w", err)
	}

	utils.Logger.Info("Успешное подключение к базе данных")
	return nil
}

// Close закрывает соединение с базой данных
func Close() error {
	if DbPostgres != nil {
		if err := DbPostgres.Close(); err != nil {
			return err
		}
	}
	return nil
}
