package v1

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	password "github.com/vzglad-smerti/password_hash"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Credentials struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// generateToken создает новый JWT токен с данными пользователя и временем истечения
func generateToken(nickname string) (string, error) {
	expirationTime := time.Now().Add(20 * time.Hour)
	claims := &Claims{
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(jwtKey)
	fmt.Println("\n\n", ss)
	return token.SignedString(jwtKey)
}

// @Summary User login
// @Description Login using nickname and password to generate a JWT token
// @Tags sign
// @Accept json
// @Produce json
// @Param creds body Credentials true "User credentials"
// @Success 200 {object} statusResponse "JWT token"
// @Failure 400 {object} errorResponse "Invalid request"
// @Failure 401 {object} errorResponse "Unauthorized error"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /login [post]
func login(c *gin.Context) {
	var creds Credentials
	// Привязываем данные из запроса в структуру Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// Используем GORM для получения данных пользователя
	var pc models.ProfileCheck
	if err := database.DbPostgres.Where("nickname = ?", creds.Nickname).First(&pc).Error; err != nil {
		// Если не найдено, то возвращаем ошибку
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Bad check profile"})
		return
	}

	// Проверяем, совпадает ли введенный пароль с сохраненным хешом
	if ok, err := password.Verify(pc.HashPassword, creds.Password); !ok || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Password error!!!"})
		return
	}

	// Генерируем токен
	token, err := generateToken(creds.Nickname)
	if err != nil {
		utils.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	// Отправляем ответ с токеном
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// authMiddleware - middleware для проверки валидности JWT токена
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем JWT токен из заголовка
		tokenString := strings.Split(c.GetHeader("Authorization"), " ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Инициализируем структуру для хранения данных токена
		claims := &Claims{}
		// Пытаемся распарсить токен
		token, err := jwt.ParseWithClaims(tokenString[1], claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// Если токен невалиден или произошла ошибка, отклоняем запрос
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		// Если токен валиден, передаем выполнение дальше
		c.Next()
	}
}
