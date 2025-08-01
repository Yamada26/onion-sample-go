package presentation_test

import (
	"net/http"
	"net/http/httptest"
	"onion-sample-go/domain"
	"onion-sample-go/presentation"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

type StubArticleUsecase struct {
	GetArticleByIdFunc func(id int) (*domain.Article, error)
	CreateArticleFunc  func(article *domain.Article) (*domain.Article, error)
}

func (au *StubArticleUsecase) CreateArticle(article *domain.Article) (*domain.Article, error) {
	return au.CreateArticleFunc(article)
}

func (au *StubArticleUsecase) GetArticleById(id int) (*domain.Article, error) {
	return au.GetArticleByIdFunc(id)
}

func TestGetArticleById_Success(t *testing.T) {
	mockUsecase := &StubArticleUsecase{
		GetArticleByIdFunc: func(id int) (*domain.Article, error) {
			return domain.NewArticle(id, "Test Article")
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
