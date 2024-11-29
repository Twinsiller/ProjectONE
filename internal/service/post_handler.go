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

var posts = []models.Post{}

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

	utils.Logger.Printf("%v", profiles)

	c.JSON(http.StatusOK, profiles)
	posts = []models.Post{}
}

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

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	p := models.Post{}

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(post_handler.go|UpdatePost|)|:", err)
		return
	}

	_, err := database.DbPostgres.Exec("UPDATE posts SET title = $1, description = $2, date_last_modified = $3  WHERE id = $4",
		p.Title, p.Description, time.Now(), p.Likes, id,
	)
	if err != nil {
		utils.Logger.Panic("Insert isn't done(post_handler.go|UpdatePost|):", err)
		return
	}
	//fmt.Println(result.LastInsertId()) // не поддерживается
	//fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusAccepted, p)
}

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
