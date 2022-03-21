package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a ArticleTag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&ArticleTag{}).Where("article_id = ?", a.ArticleID).Update(values).Error
}

func (a ArticleTag) Delete(db *gorm.DB, articleId uint32) error {
	var articleTags []*ArticleTag
	return db.Model(&ArticleTag{}).Where("article_id = ?", articleId).Delete(&articleTags).Error
}
