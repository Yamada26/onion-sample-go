package presentation

import (
	"net/http"
	"onion-sample-go/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleUsecase interface {
	GetArticleById(id int) (*domain.Article, error)
}

type ArticleHandler struct {
	articleUsecase ArticleUsecase
}

func NewArticleHandler(au ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{articleUsecase: au}
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
