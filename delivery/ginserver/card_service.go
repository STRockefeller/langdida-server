package ginserver

import (
	"net/http"

	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/service/instance"
	"github.com/gin-gonic/gin"
)

func setupCardService(router *gin.Engine, service *instance.CardService) {
	router.POST("/card/create", newCreateCardHandler(service))
	router.POST("/card/edit", newEditCardHandler(service))
	router.GET("/card/get", newGetCardHandler(service))
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

func newEditCardHandler(service *instance.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var card protomodels.Card
		if err := ctx.BindJSON(&card); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		service.EditCard(ctx, card)
		ctx.JSON(http.StatusOK, "OK")
	}
}

func langMapping(lang string) protomodels.Language {
	switch lang {
	case "en":
		return protomodels.Language_ENGLISH
	case "jp":
		return protomodels.Language_JAPANESE
	case "fr":
		return protomodels.Language_FRENCH
	}
	return protomodels.Language_ENGLISH
}

func newGetCardHandler(service *instance.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		lang := ctx.Query("language")
		word := ctx.Query("word")
		card, err := service.GetCard(ctx, protomodels.CardIndex{Language: langMapping(lang), Name: word})
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, card)
	}
}
