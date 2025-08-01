package usecase

import "onion-sample-go/domain"

type ArticleUsecase struct {
	articleRepository domain.ArticleRepository
}

func NewArticleUsecase(articleRepository domain.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{articleRepository: articleRepository}
}

func (au *ArticleUsecase) GetArticleByID(id int) (*domain.Article, error) {
	article, err := au.articleRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}
