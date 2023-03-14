package ginserver

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Run(port int) {
	router := gin.Default()
	router.Run(":" + strconv.Itoa(port))
	router.POST("/card/create", createCardHandler)
}

func createCardHandler(c *gin.Context) {
	fmt.Println("not implement")
}
