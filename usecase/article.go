package usecase

import "onion-sample-go/domain"

type ArticleUsecase struct {
	articleRepository domain.ArticleRepository
}

func NewArticleUsecase(ar domain.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{articleRepository: ar}
}

type CreateArticleDto struct {
	Id    int
	Title string
}

func (au *ArticleUsecase) CreateArticle(article *domain.Article) (*CreateArticleDto, error) {
	createdArticle, _ := au.articleRepository.Create(article)

	return &CreateArticleDto{
		Id:    createdArticle.GetId(),
		Title: createdArticle.GetTitle(),
	}, nil
}

type GetArticleByIdDto struct {
	Id    int
	Title string
}

func (au *ArticleUsecase) GetArticleById(id int) (*GetArticleByIdDto, error) {
	article, _ := au.articleRepository.FindById(id)

	return &GetArticleByIdDto{
		Id:    article.GetId(),
		Title: article.GetTitle(),
	}, nil
}
