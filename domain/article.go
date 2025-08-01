package domain

import "errors"

type Article struct {
	id    int
	title string
}

func (a *Article) GetId() int {
	return a.id
}

func (a *Article) GetTitle() string {
	return a.title
}

func NewArticle(id int, title string) (*Article, error) {
	if id <= 0 {
		return nil, errors.New("id must be a non-negative integer")
	}

	if title == "" {
		return nil, errors.New("title must not be empty")
	}

	return &Article{id: id, title: title}, nil
}

type ArticleRepository interface {
	FindById(id int) (article *Article, err error)
}
