package service

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"net/http"
	"strconv"

	password "github.com/vzglad-smerti/password_hash"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var profiles = []models.Profile{}

// @Summary		Get profiles
// @Security		ApiKeyAuth
// @Description	Retrieve a list of profiles for a specific account by account ID with pagination
// @Tags			authors
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of profiles per page (default: 5)"
// @Success		200		{array}		models.Profile
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/profiles [get]
func GetProfiles(c *gin.Context) {
	// Получение параметров из запроса
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))   // Номер страницы, по умолчанию 1
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5")) // Количество элементов на странице, по умолчанию 5

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit // Вычисление смещения

	var profiles []models.Profile

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Find(&profiles).Error
	if err != nil {
		utils.Logger.Panic(err)
		return
	}

	//utils.Logger.Printf("%v", profiles)

	c.JSON(http.StatusOK, profiles)
}

// @Summary		Get profile by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific profile by its ID
// @Tags			authors
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Account ID"
// @Success		200	{object}	models.Profile
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/profiles/{id} [get]
func GetProfileById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var profile models.Profile
	err := database.DbPostgres.First(&profile, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "profile not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Panic("Неудачный запрос|(profile_handler.go|GetProfileById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, profile)
}

// @Summary		Create a new profile
// @Description	Creates a new profile by accepting profile details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			profile	body		models.Profile	true	"Profile data"
// @Success		201	{object}	models.Profile
// @Failure		400	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/register [post]
func CreateProfile(c *gin.Context) {
	p := models.Profile{}

	// Парсим JSON из тела запроса в структуру Profile
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(profile_handler.go|CreateProfile|)|:", err)
		return
	}

	// Хеширование пароля
	if hash, err := password.Hash(p.HashPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Problem with password hashing"})
		utils.Logger.Panic("Hash wasn't working(profile_handler.go|CreateProfile|):", err)
		return
	} else {
		p.HashPassword = hash
	}

	// Создаем новый профиль в базе данных с использованием GORM
	if err := database.DbPostgres.Create(&p).Error; err != nil {
		utils.Logger.Panic("Insert isn't done(profile_handler.go|CreateProfile|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Отправляем успешный ответ с созданным профилем
	c.JSON(http.StatusCreated, p)
}

// @Summary		Update an existing profile
// @Security		ApiKeyAuth
// @Description	Update an existing profile's information by profile ID
// @Tags			authors
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"Profile ID"
// @Param			profile	body		models.Profile	true	"Updated profile data"
// @Success		202	{object}	models.Profile
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/profiles/{id} [put]
func UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	p := models.Profile{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(profile_handler.go|UpdateProfile|)|:", err)
		return
	}

	// Хеширование пароля
	if hash, err := password.Hash(p.HashPassword); err != nil {
		utils.Logger.Panic("Hash wasn't working(profile_handler.go|UpdateProfile|):", err)
		return
	} else {
		p.HashPassword = hash
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.Profile{}).Where("id = ?", id).Updates(p).Error; err != nil {
		utils.Logger.Panic("Update isn't done(profile_handler.go|UpdateProfile|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, p)
}

// @Summary		Delete a profile by ID
// @Security		ApiKeyAuth
// @Description	Delete a profile from the system by its ID
// @Tags			authors
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Profile ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/profiles/{id} [delete]
func DeleteProfile(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.Profile{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profile wasn't deleted"})
		utils.Logger.Error("Delete isn't done(profile_handler.go|DeleteProfile|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "profile was deleted"})
}
