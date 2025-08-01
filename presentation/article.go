package presentation

import (
	"net/http"
	"onion-sample-go/domain"
	"onion-sample-go/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleUsecase interface {
	GetArticleById(id int) (*usecase.GetArticleByIdDto, error)
	CreateArticle(article *domain.Article) (*usecase.CreateArticleDto, error)
}

type ArticleHandler struct {
	articleUsecase ArticleUsecase
}

func NewArticleHandler(au ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{articleUsecase: au}
}

func (ah *ArticleHandler) CreateArticle(ctx *gin.Context) {
	var reqBody struct {
		Title string `json:"title" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	articleToCreate, err := domain.NewArticle(0, reqBody.Title)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdArticle, err := ah.articleUsecase.CreateArticle(articleToCreate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":    createdArticle.Id,
		"title": createdArticle.Title,
	})
}

func (ah *ArticleHandler) GetArticleById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	article, err := ah.articleUsecase.GetArticleById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    article.Id,
		"title": article.Title,
	})
}
