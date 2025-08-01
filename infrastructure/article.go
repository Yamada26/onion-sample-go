package infrastructure

import "onion-sample-go/domain"

type ArticleRepository struct{}

func (ar *ArticleRepository) FindByID(id int) (*domain.Article, error) {
	article, _ := domain.NewArticle(id, "hoge")
	return article, nil
}
