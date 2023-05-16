package ginserver

import (
	"github.com/STRockefeller/langdida-server/service/instance"
	"github.com/gin-gonic/gin"
)

func setupIOService(router *gin.Engine, service *instance.IOService) {
	router.POST("/io/import/file", newImportFileHandler(service))
	router.POST("/io/import/url", newImportURLHandler(service))
}

// todo: fix: should import from client
func newImportFileHandler(service *instance.IOService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		filePath := "/tmp/" + file.Filename
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		content, err := service.ImportFromFile(ctx, filePath)
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		ctx.JSON(200, gin.H{"content": content})
	}
}

func newImportURLHandler(service *instance.IOService) func(*gin.Context) {
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
