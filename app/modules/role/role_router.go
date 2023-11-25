package role

import (
	"api-gin/app/modules/role/infrastructure"
	"fmt"

	"github.com/gin-gonic/gin"
)

var roleController = infrastructure.RoleController{}

func Router(router *gin.Engine) {
	fmt.Println("\nROUTER roles")
	router.GET("/roles", roleController.Find)
}
