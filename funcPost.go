package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Author      Profile `json:"author"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Likes       int     `json:"likes"`
}

var countIdPosts = 3
var posts = []Post{
	{ID: 0, Title: "Holiday", Author: profiles[0],
		Date: "2 Jan 2006 15:04:05", Description: "It was a great holiday", Likes: 0},
	{ID: 1, Title: "Sunday", Author: profiles[2],
		Date: "3 Jan 2012 19:10:23", Description: "I enjoy my life", Likes: 2},
	{ID: 2, Title: "Monday", Author: profiles[0],
		Date: "5 Jan 2024 23:04:40", Description: "AGAIN!!!!", Likes: 304},
}

func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	for _, post := range posts {
		if post.ID == id {
			c.JSON(http.StatusOK, post)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

func CreatePost(c *gin.Context) {
	var newPost Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	fmt.Println(newPost)
	for _, author := range profiles {
		if newPost.Author.ID == author.ID {
			newPost.ID = countIdPosts
			countIdPosts++
			newPost.Date = time.Now().Format("2 Jan 2006 15:04:05")
			newPost.Likes = 0
			posts = append(posts, newPost)
			c.JSON(http.StatusCreated, newPost)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "not found author"})
}

func UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
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

func DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "post was deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "post not found for delete"})
}
