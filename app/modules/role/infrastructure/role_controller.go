package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

var roleService = RoleService{}

func (uc *RoleController) Find(c *gin.Context) {
	var roles = roleService.find()

	c.IndentedJSON(http.StatusOK, roles)
}
