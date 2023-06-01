package ginserver

import (
	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/service"
	"github.com/gin-gonic/gin"
)

func setupExerciseService(router *gin.Engine, service service.ExerciseService) {
	router.POST("/exercise/choice", newChoiceProblemsHandler(service))
	router.POST("/exercise/filling", newFillingProblemsHandler(service))
}

func newChoiceProblemsHandler(service service.ExerciseService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var cards []protomodels.CardIndex
		if err := ctx.BindJSON(&cards); err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		problems, answers, err := service.CreateChoiceProblems(ctx, cards)
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		ctx.JSON(200, gin.H{"problems": problems, "answers": answers})
	}
}

func newFillingProblemsHandler(service service.ExerciseService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var cards []protomodels.CardIndex
		if err := ctx.BindJSON(&cards); err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		problems, answers, err := service.CreateFillingProblems(ctx, cards)
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		ctx.JSON(200, gin.H{"problems": problems, "answers": answers})
	}
}
