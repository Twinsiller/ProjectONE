package service

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var posts = []models.Post{}

// @Summary		Get all posts
// @Description	Retrieve a list of all posts in the system
// @Tags			posts
// @Accept			json
// @Produce		json
// @Success		200	{array}	models.Post
// @Failure		500	{object}	errorResponse
// @Router			/v1/posts [get]
func GetPosts(c *gin.Context) {
	rows, err := database.DbPostgres.Query("select id, id_author, title, description, date_publication, date_last_modified, likes from posts")
	if err != nil {
		utils.Logger.Panic(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := models.Post{}
		err := rows.Scan(&p.Id, &p.IdAuthor, &p.Title, &p.Description, &p.DatePublication, &p.DateLastModified, &p.Likes)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	utils.Logger.Printf("%v", posts)

	c.JSON(http.StatusOK, posts)
	posts = []models.Post{}
}

// @Summary		Get a post by ID
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
	//utils.Logger.Info("GetProfileByID is working\n(post_handler.go|GetPostByID|):\n")
	id := c.Param("id")

	row := database.DbPostgres.QueryRow("select * from posts where id = $1", id)
	p := models.Post{}

	err := row.Scan(&p.Id, &p.IdAuthor, &p.Title, &p.Description, &p.DatePublication, &p.DateLastModified, &p.Likes)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
		utils.Logger.Panic("Not correct request|(post_handler.go|GetPostByID|)|:", err)
	}

	//utils.Logger.Info("Request is done\n(post_handler.go|GetPostByID|):\n")
	c.JSON(http.StatusOK, p)
}

// @Summary		Create a new post
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
	p := models.Post{}

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(post_handler.go|CreatePost|)|:", err)
		return
	}

	_, err := database.DbPostgres.Exec("insert into posts (id_author, title, description) values ( $1, $2, $3)",
		p.IdAuthor, p.Title, p.Description,
	)
	if err != nil {
		utils.Logger.Panic("Insert isn't done(post_handler.go|CreatePost|):", err)
		return
	}
	// fmt.Println(result.LastInsertId()) // не поддерживается (Из-за Postgres)
	// fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusCreated, p)
}

// @Summary		Update an existing post
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
	p := models.Post{}

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(post_handler.go|UpdatePost|)|:", err)
		return
	}

	_, err := database.DbPostgres.Exec("UPDATE posts SET title = $1, description = $2, date_last_modified = now()  WHERE id = $3",
		p.Title, p.Description, id,
	)
	if err != nil {
		utils.Logger.Panic("Insert isn't done(post_handler.go|UpdatePost|):", err)
		return
	}
	//fmt.Println(result.LastInsertId()) // не поддерживается
	//fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusAccepted, p)
}

// @Summary		Delete a post by ID
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

	_, err := database.DbPostgres.Exec("delete from posts where id = $1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "post wasn't deleted"})
		utils.Logger.Panic("Insert isn't done(post_handler.go|DeletePost|):", err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "post was deleted"})
}
