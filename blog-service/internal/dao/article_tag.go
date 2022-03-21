package dao

import "github.com/ydnju/go_tour/blog-service/internal/model"

func (d *Dao) CreateArticleTag(articleId uint32, tagId uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		TagID:     tagId,
		ArticleID: articleId,
		Model:     &model.Model{CreatedBy: createdBy},
	}

	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleId uint32, tagId uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{
		TagID:     tagId,
		ArticleID: articleId,
		Model:     &model.Model{ModifiedBy: modifiedBy},
	}

	values := map[string]interface{}{
		"tag_id":      tagId,
		"modified_by": modifiedBy,
	}
	return articleTag.Update(d.engine, values)
}

func (d *Dao) DeleteArticleTag(articleId uint32) error {
	articleTag := model.ArticleTag{
		ArticleID: articleId,
	}

	return articleTag.Delete(d.engine, articleId)
}
