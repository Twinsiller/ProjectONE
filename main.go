package main

import (
	"github.com/gin-gonic/gin"
)

/*func main() {
	// Создаем новый роутер Gin
	router := gin.Default()

	// Определяем маршруты
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/uou", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"textoooo": "opacha",
		})
	})

	// Запускаем сервер на порту 8080
	router.Run(":8080")
}*/

/*type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Likes       string `json:"likes"`
}*/
/*
var posts = []Post{
	{ID: "1", Title: "Holiday", Author: "Twinsiller",
		Date: "19.10.2024", Description: "It was a great holiday", Likes: "0"},
	{ID: "2", Title: "Sunday", Author: "Twinsiller",
		Date: "20.10.2024", Description: "I enjoy my life", Likes: "2"},
	{ID: "3", Title: "Monday", Author: "Muha",
		Date: "21.10.2024", Description: "AGAIN!!!!", Likes: "304"},
}

func main() {
	router := gin.Default()

	// Получение всех книг
	router.GET("/books", getPosts)

	// Получение книги по ID
	router.GET("/books/:id", getPostByID)

	// Создание новой книги
	router.POST("/books", createPost)

	// Обновление существующей книги
	router.PUT("/books/:id", updatePost)

	// Удаление книги
	router.DELETE("/books/:id", deletePost)

	router.Run(":8080")
}

func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func getPostByID(c *gin.Context) {
	id := c.Param("id")
	for _, post := range posts {
		if post.ID == id {
			c.JSON(http.StatusOK, post)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

func createPost(c *gin.Context) {
	var newPost Post

	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	posts = append(posts, newPost)
	c.JSON(http.StatusCreated, newPost)
}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	var updatedPost Post

	if err := c.BindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, post := range posts {
		if post.ID == id {
			posts[i] = updatedPost
			c.JSON(http.StatusOK, updatedPost)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "post was deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "post not found for delete"})
}*/

func main() {
	router := gin.Default()

	// Получение всех профилей
	router.GET("/profiles", GetProfiles)
	// Получение всех постов
	router.GET("/posts", GetPosts)

	// Получение поста по ID
	router.GET("/profiles/:id", GetProfileByID)
	// Получение профиля по ID
	router.GET("/posts/:id", GetPostByID)

	// Создание нового профиля
	router.POST("/profiles", CreateProfile)
	// Создание новой поста
	router.POST("/posts", CreatePost)

	// Обновление существующего профиля
	router.PUT("/profiles/:id", UpdateProfile)
	// Обновление существующего поста
	router.PUT("/posts/:id", UpdatePost)

	// Удаление профиля
	router.DELETE("/profiles/:id", DeleteProfile)
	// Удаление поста
	router.DELETE("/posts/:id", DeletePost)

	router.Run(":8080")
}
