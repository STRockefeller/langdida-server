package ginserver

import (
	"net/http"
	"strconv"

	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/storage"
	"github.com/gin-gonic/gin"
)

func Run(port int, storage storage.Storage) {
	router := gin.Default()
	router.Run(":" + strconv.Itoa(port))
	router.POST("/card/create", newCreateCardHandler(storage))
}

func newCreateCardHandler(storage storage.Storage) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var card protomodels.Card
		if err := ctx.BindJSON(&card); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		storage.CreateCard(ctx, card)
		ctx.JSON(http.StatusOK, "OK")
	}
}
