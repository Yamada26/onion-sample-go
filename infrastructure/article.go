package infrastructure

import (
	"onion-sample-go/domain"
)

type ArticleModel struct {
	ID    int    `gorm:"primaryKey;column:id"`
	Title string `gorm:"column:title"`
}

func (ArticleModel) TableName() string {
	return "articles"
}

type ArticleRepository struct{}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

/*
指定された ID の Article を取得する。
存在しない場合は、gorm.ErrRecordNotFound を返す。
*/
func (ar *ArticleRepository) FindById(id int) (*domain.Article, error) {
	db := GetDB()
	var model ArticleModel
	if err := db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return domain.NewArticle(model.ID, model.Title)
}
