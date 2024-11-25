package v1

import "github.com/gin-gonic/gin"

func Handler() {
	Router := gin.Default()

	Router.Run(":8080")
}
