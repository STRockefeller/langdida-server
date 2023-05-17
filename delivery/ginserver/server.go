package ginserver

import (
	"strconv"

	"github.com/STRockefeller/langdida-server/service/instance"
	"github.com/STRockefeller/langdida-server/storage"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run(port int, storage storage.Storage) {
	router := gin.Default()
	router.Run(":" + strconv.Itoa(port))
	zap.L().Info("server started", zap.Int("port", port))

	setupCardService(router, instance.NewCardService(storage))
	setupExerciseService(router, instance.NewExerciseService(storage))
	setupIOService(router, instance.NewIOService())
	router.GET("/ping", pingHandler)
}

func pingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}
