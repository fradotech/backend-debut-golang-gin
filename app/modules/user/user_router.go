package user

import (
	"api-gin/app/modules/user/infrastructure"
	"fmt"

	"github.com/gin-gonic/gin"
)

var userController = infrastructure.UserController{}

func Router(router *gin.Engine) {
	fmt.Println("\nROUTER users")
	router.GET("/users", userController.Find)
}
