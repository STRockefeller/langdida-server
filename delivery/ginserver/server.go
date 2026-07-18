package ginserver

import (
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/STRockefeller/langdida-server/service/instance"
	"github.com/STRockefeller/langdida-server/storage"
)

func Run(port int, storage storage.Storage) {
	router := gin.Default()

	router.Use(cors.New(localCORSConfig()))

	setupCardService(router, instance.NewCardService(storage))
	setupExerciseService(router, instance.NewExerciseService(storage))
	setupIOService(router, instance.NewIOService())
	router.GET("/ping", pingHandler)

	address := net.JoinHostPort("127.0.0.1", strconv.Itoa(port))
	if err := router.Run(address); err != nil {
		zap.L().Fatal(err.Error())
	}
	zap.L().Info("server started", zap.String("address", address))
}

func localCORSConfig() cors.Config {
	return cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		MaxAge:       12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			parsedOrigin, err := url.Parse(origin)
			if err != nil || (parsedOrigin.Scheme != "http" && parsedOrigin.Scheme != "https") {
				return false
			}
			host := parsedOrigin.Hostname()
			return host == "localhost" || net.ParseIP(host).IsLoopback()
		},
	}
}

func pingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}
