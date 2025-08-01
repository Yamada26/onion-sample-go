package infrastructure

import "onion-sample-go/domain"

type ArticleRepository struct{}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (ar *ArticleRepository) FindById(id int) (*domain.Article, error) {
	article, _ := domain.NewArticle(id, "hoge")
	return article, nil
}
