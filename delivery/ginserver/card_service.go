package ginserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"

	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/service"
	"github.com/STRockefeller/langdida-server/storage"
)

func setupCardService(router *gin.Engine, service service.CardService) {
	router.POST("/card/create", newCreateCardHandler(service))
	router.PUT("/card/edit", newEditCardHandler(service))
	router.GET("/card/get", newGetCardHandler(service))
	router.GET("/card/dictionary/meanings", newSearchMeaningsHandler(service))
	router.GET("/card/list", newListCardsHandler(service))
	router.GET("/card/index/list", newListCardIndicesHandler(service))
	router.GET("/card/association", newGetAssociationHandler(service))
	router.POST("/card/association/create", newCreateAssociationHandler(service))
}

func newCreateCardHandler(service service.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var card protomodels.Card
		if err := ctx.ShouldBindJSON(&card); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := service.CreateCard(ctx, card); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
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

func newListCardsHandler(service service.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		needReview := ctx.Query("needReview")
		language := ctx.Query("language")
		label := ctx.Query("label")

		mappedLanguage := protomodels.LangMapping(language)

		cards, err := service.ListCards(ctx, storage.ListCardsConditions{
			NeedReview: lo.Ternary(needReview == "true", true, false),
			Language:   lo.Ternary(language != "", &mappedLanguage, nil),
			Label:      lo.Ternary(label != "", label, ""),
		})

		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, cards)
	}
}

func newListCardIndicesHandler(service service.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		indices, err := service.ListIndices(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, indices)
	}
}

func newGetAssociationHandler(service service.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		lang := ctx.Query("language")
		word := ctx.Query("word")
		cards, err := service.GetAssociations(ctx, protomodels.CardIndex{Language: protomodels.LangMapping(lang), Name: word})
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, cards)
	}
}

func newCreateAssociationHandler(service service.CardService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var conditions storage.CreateAssociationConditions
		if err := ctx.ShouldBindJSON(&conditions); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := service.CreateAssociations(ctx, conditions); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, "OK")
	}
}
