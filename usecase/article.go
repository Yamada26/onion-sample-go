package usecase

import "onion-sample-go/domain"

type ArticleUsecase struct {
	articleRepository domain.ArticleRepository
}

func NewArticleUsecase(ar domain.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{articleRepository: ar}
}

func (au *ArticleUsecase) GetArticleById(id int) (*domain.Article, error) {
	return au.articleRepository.FindById(id)
}
