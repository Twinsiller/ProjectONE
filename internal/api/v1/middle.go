package v1

import (
	database "ProjectONE/internal/database/postgres"
	"ProjectONE/internal/models"
	"ProjectONE/pkg/utils"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	password "github.com/vzglad-smerti/password_hash"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Credentials struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

func generateToken(nickname string) (string, error) {
	expirationTime := time.Now().Add(20 * time.Hour)
	claims := &Claims{
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(jwtKey)
	fmt.Println("\n\n", ss)
	return token.SignedString(jwtKey)
}

func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	row := database.DbPostgres.QueryRow("select nickname, hash_password from authors where nickname = $1", creds.Nickname)
	pc := models.ProfileCheck{}
	fmt.Println(row)
	err := row.Scan(&pc.Nickname, &pc.HashPassword)
	fmt.Println("\n\n", creds.Nickname, creds.Password)
	fmt.Println(pc.Nickname, pc.HashPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Bad check profile"})
		return
	}
	if creds.Nickname != pc.Nickname {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User doesn't exist"})
		return
	}
	if ok, err := password.Verify(pc.HashPassword, creds.Password); !ok || err != nil {
		fmt.Println("ok = ", ok)
		fmt.Println("error = ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Password error!!!"})
		return
	}

	token, err := generateToken(creds.Nickname)
	if err != nil {
		utils.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.GetHeader("Authorization"), " ")
		//fmt.Println("\n\n", tokenString)
		claims := &Claims{}
		//fmt.Println("\n\nclaims.Username ", claims.Username)
		token, err := jwt.ParseWithClaims(tokenString[1], claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
		//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		c.Next()
	}
}
