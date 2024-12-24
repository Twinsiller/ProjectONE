package service

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var comments = []models.Comment{}

// @Summary      Get all comments
// @Security		ApiKeyAuth
// @Description  Retrieve a list of all comments from the database
// @Tags         comments
// @Produce      json
// @Success      200 {array} models.Comment
// @Failure      500 {object} errorResponse
// @Router       /v1/comments [get]
func GetComments(c *gin.Context) {
	rows, err := database.DbPostgres.Query("select * from comments")
	fmt.Println(rows)
	if err != nil {
		utils.Logger.Panic(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		cm := models.Comment{}
		err := rows.Scan(&cm.Id, &cm.IdAuthor, &cm.Text, &cm.DatePublication, &cm.DatePublication, &cm.IdPost)
		fmt.Println(cm)
		if err != nil {
			fmt.Println(err)
			continue
		}
		comments = append(comments, cm)
	}

	//utils.Logger.Printf("%v", comments)

	c.JSON(http.StatusOK, comments)
	comments = []models.Comment{}
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
	//utils.Logger.Info("GetCommentById is working\n(comment_handler.go|GetCommentById|):\n")
	id := c.Param("id")
	row := database.DbPostgres.QueryRow("select * from comments where id = $1", id)
	cm := models.Comment{}
	err := row.Scan(&cm.Id, &cm.IdAuthor, &cm.Text, &cm.DatePublication, &cm.DatePublication, &cm.IdPost)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "comment not found"})
		utils.Logger.Panic("Not correct request|(comment_handler.go|GetCommentById|)|:", err)
	}
	//utils.Logger.Info("Request is done(comment_handler.go|GetCommentById|):")
	c.JSON(http.StatusOK, cm)
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
	cm := models.Comment{}

	if err := c.BindJSON(&cm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(comment_handler.go|CreateComment|)|:", err)
		return
	}

	_, err := database.DbPostgres.Exec("insert into comments (id_author, id_post, text_comment) values ( $1, $2, $3)",
		cm.IdAuthor, cm.IdPost, cm.Text,
	)
	if err != nil {
		utils.Logger.Panic("Insert isn't done(comment_handler.go|CreateComment|):", err)
		return
	}
	// fmt.Println(result.LastInsertId()) // не поддерживается (Из-за Postgres)
	// fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusCreated, cm)
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
	cm := models.Comment{}

	if err := c.BindJSON(&cm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(comment_handler.go|UpdateComment|)|:", err)
		return
	}

	_, err := database.DbPostgres.Exec("UPDATE comments SET id_author = $1, id_post = $2, text_comment = $3, date_last_modified = $4 WHERE id = $7",
		cm.IdAuthor, cm.IdPost, cm.Text, time.Now(), id,
	)
	if err != nil {
		utils.Logger.Panic("Insert isn't done(comment_handler.go|UpdateComment|):", err)
		return
	}
	//fmt.Println(result.LastInsertId()) // не поддерживается
	//fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusAccepted, cm)
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

	_, err := database.DbPostgres.Exec("delete from comments where id = $1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profile wasn't deleted"})
		utils.Logger.Panic("Insert isn't done(comment_handler.go|DeleteComment|):", err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "comment was deleted"})
}
