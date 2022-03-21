package dao

import (
	"github.com/ydnju/go_tour/blog-service/internal/model"
)

func (d *Dao) GetArticle(id uint32, state uint8) (*model.Article, error) {
	article := model.Article{
		State: state,
		Model: &model.Model{ID: id},
	}

	return article.Get(d.engine)
}

func (d *Dao) GetArticleList(tagId uint32, state uint8, pageOffset, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{
		State: state,
	}

	return article.ListByTagId(d.engine, tagId, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(tagId uint32, title string, desc string, content string,
	coverImageUrl string, state uint8, createdBy string) (*model.Article, error) {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
		Model:         &model.Model{CreatedBy: createdBy},
	}

	ar, err := article.Create(d.engine)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (d *Dao) UpdateArticle(
	id uint32, title string, desc string, content string,
	coverImageUrl string, state uint8, modifiedBy string) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
		Model:         &model.Model{ID: id, ModifiedBy: modifiedBy},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if title != "" {
		values["title"] = title
	}
	if desc != "" {
		values["desc"] = desc
	}
	if content != "" {
		values["content"] = content
	}
	if coverImageUrl != "" {
		values["coverImageUrl"] = coverImageUrl
	}

	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}
