package presentation

import (
	"onion-sample-go/infrastructure"
	"onion-sample-go/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	articleRepository := infrastructure.NewArticleRepository()
	articleUsecase := usecase.NewArticleUsecase(articleRepository)
	articleHandler := NewArticleHandler(articleUsecase)
	router.GET("/articles/:id", articleHandler.GetArticleById)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "successful"})
	})

	return router
}
