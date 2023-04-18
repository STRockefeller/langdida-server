package ginserver

import (
	"net/http"
	"strconv"

	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/service/instance"
	"github.com/STRockefeller/langdida-server/storage"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run(port int, storage storage.Storage) {
	router := gin.Default()
	router.Run(":" + strconv.Itoa(port))
	zap.L().Info("server started", zap.Int("port", port))
	router.POST("/card/create", newCreateCardHandler(instance.NewCardService(storage)))
}

func newCreateCardHandler(service *instance.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var card protomodels.Card
		if err := ctx.BindJSON(&card); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		service.CreateCard(ctx, card)
		ctx.JSON(http.StatusOK, "OK")
	}
}
