package service

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"fmt"
	"net/http"

	password "github.com/vzglad-smerti/password_hash"

	"github.com/gin-gonic/gin"
)

var profiles = []models.Profile{}

func GetProfiles(c *gin.Context) {
	rows, err := database.DbPostgres.Query("select * from authors")
	if err != nil {
		utils.Logger.Panic(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := models.Profile{}
		err := rows.Scan(&p.Id, &p.Nickname, &p.HashPassword, &p.Status, &p.AccessLevel, &p.Firstname, &p.Lastname, &p.CreatedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		profiles = append(profiles, p)
	}

	//utils.Logger.Printf("%v", profiles)

	c.JSON(http.StatusOK, profiles)
	profiles = []models.Profile{}
}

func GetProfileById(c *gin.Context) {
	//utils.Logger.Info("GetProfileByID is working\n(profile_handler.go|GetProfileByID|):\n")
	id := c.Param("id")
	row := database.DbPostgres.QueryRow("select * from authors where id = $1", id)
	p := models.Profile{}
	err := row.Scan(&p.Id, &p.Nickname, &p.HashPassword, &p.Status, &p.AccessLevel, &p.Firstname, &p.Lastname, &p.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profile not found"})
		utils.Logger.Panic("Not correct request|(profile_handler.go|GetProfileByID|)|:", err)
	}
	//utils.Logger.Info("Request is done\n(profile_handler.go|GetProfileByID|):\n")
	c.JSON(http.StatusOK, p)
}

func CreateProfile(c *gin.Context) {
	p := models.Profile{}

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(profile_handler.go|CreateProfile|)|:", err)
		return
	}

	if hash, err := password.Hash(p.HashPassword); err != nil {
		utils.Logger.Panic("Hash wasn't working(profile_handler.go|CreateProfile|):", err)
		return
	} else {
		p.HashPassword = hash
	}

	_, err := database.DbPostgres.Exec("insert into authors (nickname, hash_password, access_level, firstname, lastname) values ( $1, $2, $3, $4, $5)",
		p.Nickname, p.HashPassword, p.AccessLevel, p.Firstname, p.Lastname,
	)
	if err != nil {
		utils.Logger.Panic("Insert isn't done(profile_handler.go|CreateProfile|):", err)
		return
	}
	// fmt.Println(result.LastInsertId()) // не поддерживается (Из-за Postgres)
	// fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusCreated, p)
}

func UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	p := models.Profile{}

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Panic("Data is bad|(profile_handler.go|UpdateProfile|)|:", err)
		return
	}

	if hash, err := password.Hash(p.HashPassword); err != nil {
		utils.Logger.Panic("Hash wasn't working(profile_handler.go|UpdateProfile|):", err)
		return
	} else {
		p.HashPassword = hash
	}

	_, err := database.DbPostgres.Exec("UPDATE authors SET nickname = $1, hash_password = $2, status = $3, access_level = $4, firstname = $5, lastname = $6  WHERE id = $7",
		p.Nickname, p.HashPassword, p.Status, p.AccessLevel, p.Firstname, p.Lastname, id,
	)
	if err != nil {
		utils.Logger.Panic("Insert isn't done(profile_handler.go|UpdateProfile|):", err)
		return
	}
	//fmt.Println(result.LastInsertId()) // не поддерживается
	//fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusAccepted, p)
}

func DeleteProfile(c *gin.Context) {
	id := c.Param("id")

	_, err := database.DbPostgres.Exec("delete from authors where id = $1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profile wasn't deleted"})
		utils.Logger.Panic("Insert isn't done(profile_handler.go|UpdateProfile|):", err)
		return

	}

	c.JSON(http.StatusAccepted, gin.H{"message": "profile was deleted"})
}
