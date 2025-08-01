package presentation_test

import (
	"net/http"
	"net/http/httptest"
	"onion-sample-go/domain"
	"onion-sample-go/presentation"
	"onion-sample-go/usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

type StubArticleUsecase struct {
	CreateArticleFunc  func(article *domain.Article) (*usecase.CreateArticleDto, error)
	GetArticleByIdFunc func(id int) (*usecase.GetArticleByIdDto, error)
}

func (au *StubArticleUsecase) CreateArticle(article *domain.Article) (*usecase.CreateArticleDto, error) {
	return au.CreateArticleFunc(article)
}

func (au *StubArticleUsecase) GetArticleById(id int) (*usecase.GetArticleByIdDto, error) {
	return au.GetArticleByIdFunc(id)
}

func TestGetArticleById_Success(t *testing.T) {
	mockUsecase := &StubArticleUsecase{
		GetArticleByIdFunc: func(id int) (*usecase.GetArticleByIdDto, error) {
			article, _ := domain.NewArticle(id, "Test Article")
			return &usecase.GetArticleByIdDto{
				Id:    article.GetId(),
				Title: article.GetTitle(),
			}, nil
		},
	}

	handler := presentation.NewArticleHandler(mockUsecase)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/articles", nil)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	handler.GetArticleById(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}
