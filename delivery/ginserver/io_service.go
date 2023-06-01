package ginserver

import (
	"github.com/STRockefeller/langdida-server/service"
	"github.com/gin-gonic/gin"
)

func setupIOService(router *gin.Engine, service service.IOService) {
	router.GET("/io/import/url", newImportURLHandler(service))
}

func newImportURLHandler(service service.IOService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		url := ctx.Query("url")
		content, err := service.ImportFromURL(ctx, url)
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		ctx.JSON(200, gin.H{"content": content})
	}
}
