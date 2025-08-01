package presentation

import (
	"net/http"
	"onion-sample-go/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleUsecase *usecase.ArticleUsecase
}

func NewArticleHandler(articleUsecase *usecase.ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{articleUsecase: articleUsecase}
}

func (ah *ArticleHandler) GetArticleById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	article, err := ah.articleUsecase.GetArticleById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    article.GetId(),
		"title": article.GetTitle(),
	})
}
