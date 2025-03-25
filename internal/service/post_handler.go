package service

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Get all posts
// @Security		ApiKeyAuth
// @Description	Retrieve a list of all posts in the system
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of posts per page (default: 5)"
// @Success		200	{array}	models.Post
// @Failure		500	{object}	errorResponse
// @Router			/v1/posts [get]
func GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}
	offset := (page - 1) * limit

	var posts []models.Post
	if err := database.DbPostgres.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		utils.Logger.Panic("Failed to fetch posts:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// @Summary		Get a post by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a post's details by its unique ID
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Post ID"
// @Success		200	{object}	models.Post
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/posts/{id} [get]
func GetPostById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.DbPostgres.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}
		utils.Logger.Panic("Failed to fetch post by ID:", err)
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary		Create a new post
// @Security		ApiKeyAuth
// @Description	Create a new post with title, description, and author information
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			post	body		models.Post	true	"New post data"
// @Success		201	{object}	models.Post
// @Failure		400	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/posts [post]
func CreatePost(c *gin.Context) {
	var p models.Post

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		utils.Logger.Panic("Invalid post data:", err)
		return
	}

	if err := database.DbPostgres.Create(&p).Error; err != nil {
		utils.Logger.Panic("Failed to create post:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, p)
}

// @Summary		Update an existing post
// @Security		ApiKeyAuth
// @Description	Update the details of an existing post by its ID
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Post ID"
// @Param			post	body		models.Post	true	"Updated post data"
// @Success		202	{object}	models.Post
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/posts/{id} [put]
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var p models.Post

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		utils.Logger.Panic("Invalid post data for update:", err)
		return
	}

	// Обновляем только нужные поля
	p.DateLastModified = time.Now() // Функция для текущего времени, если есть

	if err := database.DbPostgres.Model(&models.Post{}).Where("id = ?", id).Updates(p).Error; err != nil {
		utils.Logger.Panic("Failed to update post:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update post"})
		return
	}

	c.JSON(http.StatusAccepted, p)
}

// @Summary		Delete a post by ID
// @Security		ApiKeyAuth
// @Description	Delete an existing post by its unique ID
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Post ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/posts/{id} [delete]
func DeletePost(c *gin.Context) {
	id := c.Param("id")

	if err := database.DbPostgres.Delete(&models.Post{}, id).Error; err != nil {
		utils.Logger.Panic("Failed to delete post:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post was deleted"})
}
