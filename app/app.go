package main

import (
	"api-gin/app/modules/role"
	"api-gin/app/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	user.Router(router)
	role.Router(router)

	router.Run("localhost:3000")
}
