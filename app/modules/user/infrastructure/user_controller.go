package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userService = UserService{}

func (uc *UserController) Find(c *gin.Context) {
	var users = userService.find()

	c.IndentedJSON(http.StatusOK, users)
}
