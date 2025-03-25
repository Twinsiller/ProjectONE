package service

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary      Get all comments
// @Security		ApiKeyAuth
// @Description  Retrieve a list of all comments from the database
// @Tags         comments
// @Produce      json
// @Success      200 {array} models.Comment
// @Failure      500 {object} errorResponse
// @Router       /v1/comments [get]
func GetComments(c *gin.Context) {
	var comments []models.Comment
	result := database.DbPostgres.Find(&comments)
	if result.Error != nil {
		utils.Logger.Error("Failed to get comments:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// @Summary      Get comment by ID
// @Security		ApiKeyAuth
// @Description  Retrieve a specific comment by its ID from the database
// @Tags         comments
// @Produce      json
// @Param        id   path      int  true  "Comment ID"
// @Success      200  {object}  models.Comment
// @Failure      404  {object}  errorResponse
// @Router       /v1/comments/{id} [get]
func GetCommentById(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	result := database.DbPostgres.First(&comment, id)
	if result.Error != nil {
		utils.Logger.Error("Comment not found:", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"message": "comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// @Summary      Create a comment
// @Security		ApiKeyAuth
// @Description  Add a new comment to the database
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        comment  body      models.Comment  true  "Comment Data"
// @Success      201      {object}  models.Comment
// @Failure      400      {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Router       /v1/comments [post]
func CreateComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.Logger.Error("Invalid comment data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	comment.DatePublication = time.Now()

	if err := database.DbPostgres.Create(&comment).Error; err != nil {
		utils.Logger.Error("Failed to create comment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// @Summary      Update a comment
// @Security		ApiKeyAuth
// @Description  Update an existing comment's information by its ID
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id       path      int              true  "Comment ID"
// @Param        comment  body      models.Comment   true  "Updated Comment Data"
// @Success      202      {object}  models.Comment
// @Failure      400      {object}  errorResponse
// @Failure      404      {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Router       /v1/comments/{id} [put]
func UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var existingComment models.Comment

	if err := database.DbPostgres.First(&existingComment, id).Error; err != nil {
		utils.Logger.Error("Comment not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "comment not found"})
		return
	}

	var updatedComment models.Comment
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		utils.Logger.Error("Invalid update data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	existingComment.Text = updatedComment.Text
	existingComment.IdAuthor = updatedComment.IdAuthor
	existingComment.IdPost = updatedComment.IdPost
	existingComment.DateLastModified = time.Now()

	if err := database.DbPostgres.Save(&existingComment).Error; err != nil {
		utils.Logger.Error("Failed to update comment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update comment"})
		return
	}

	c.JSON(http.StatusAccepted, existingComment)
}

// @Summary      Delete a comment
// @Security		ApiKeyAuth
// @Description  Remove a comment from the database by its ID
// @Tags         comments
// @Produce      json
// @Param        id  path  int  true  "Comment ID"
// @Success      202  {object}  string
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /v1/comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	if err := database.DbPostgres.First(&comment, id).Error; err != nil {
		utils.Logger.Error("Comment not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "comment not found"})
		return
	}

	if err := database.DbPostgres.Delete(&comment).Error; err != nil {
		utils.Logger.Error("Failed to delete comment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete comment"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "comment was deleted"})
}
