package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var countIdProfiles int = 2
var profiles = []Profile{
	{ID: 0, Name: "Alex", Password: "322"},
	{ID: 1, Name: "Valerchik", Password: "52"},
	{ID: 2, Name: "Igar", Password: "894"},
}

func GetProfiles(c *gin.Context) {
	c.JSON(http.StatusOK, profiles)
}

func GetProfileByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	for _, profile := range profiles {
		if profile.ID == id {
			c.JSON(http.StatusOK, profile)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "profile not found"})
}

func CreateProfile(c *gin.Context) {
	var newProfile Profile

	if err := c.BindJSON(&newProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	countIdProfiles++
	newProfile.ID = countIdProfiles
	profiles = append(profiles, newProfile)
	c.JSON(http.StatusCreated, newProfile)
}

func UpdateProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	var updatedProfile Profile

	if err := c.BindJSON(&updatedProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, profile := range profiles {
		if profile.ID == id {
			profiles[i] = updatedProfile
			c.JSON(http.StatusOK, updatedProfile)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func DeleteProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	for i, profile := range profiles {
		if profile.ID == id {
			profiles = append(profiles[:i], profiles[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "profile was deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "profile not found for delete"})
}
