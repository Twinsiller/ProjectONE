package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Импортируем PostgreSQL драйвер
)

// DB — глобальная переменная для хранения подключения к базе данных
var DB *sql.DB

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
	return Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

// Connect устанавливает соединение с базой данных
func Connect(cfg Config) error {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Проверяем соединение
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("база данных недоступна: %w", err)
	}

	log.Println("Успешное подключение к базе данных")
	return nil
}

// Close закрывает соединение с базой данных
func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Ошибка при закрытии соединения с базой данных: %v", err)
		} else {
			log.Println("Соединение с базой данных закрыто")
		}
	}
}
