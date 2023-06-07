package ginserver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/service"
)

func setupCardService(router *gin.Engine, service service.CardService) {
	router.POST("/card/create", newCreateCardHandler(service))
	router.PUT("/card/edit", newEditCardHandler(service))
	router.GET("/card/get", newGetCardHandler(service))
	router.GET("/card/dictionary/meanings", newSearchMeaningsHandler(service))
}

func newCreateCardHandler(service service.CardService) func(*gin.Context) {
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

func newEditCardHandler(service service.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var card protomodels.Card
		if err := ctx.BindJSON(&card); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := service.EditCard(ctx, card); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, "OK")
	}
}

func newGetCardHandler(service service.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		lang := ctx.Query("language")
		word := ctx.Query("word")
		card, err := service.GetCard(ctx, protomodels.CardIndex{Language: protomodels.LangMapping(lang), Name: word})
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, card)
	}
}

func newSearchMeaningsHandler(service service.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		lang := ctx.Query("language")
		word := ctx.Query("word")
		meanings, err := service.SearchWithDictionary(ctx, protomodels.CardIndex{Language: protomodels.LangMapping(lang), Name: word})
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, meanings)
	}
}
