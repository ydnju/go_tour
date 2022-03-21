package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ydnju/go_tour/blog-service/global"
	"github.com/ydnju/go_tour/blog-service/internal/service"
	"github.com/ydnju/go_tour/blog-service/pkg/app"
	"github.com/ydnju/go_tour/blog-service/pkg/convert"
	"github.com/ydnju/go_tour/blog-service/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取文章
// @Produce  json
// @Param id path int true "文章id" minlength(3) maxlength(100)
// @Param state query int true "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/:id [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{
		ID:    convert.StrTo(c.Param("id")).MustUInt32(),
		State: 1,
	}
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
	return
}

func (a Article) List(c *gin.Context) {}

// @Summary 创建文章
// @Produce  json
// @Param tag_id body string true "标签id"
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param desc body string true "文章描述" minlength(3) maxlength(100)
// @Param content body string true "文章内容" minlength(3) maxlength(100000)
// @Param cover_image_url body string true "封面图片链接" minlength(3) maxlength(256)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app:BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新文章
// @Produce  json
// @Param id path int true "文章ID"
// @Param tag_id body string true "文章ID"
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param desc body string true "文章描述" minlength(3) maxlength(100)
// @Param content body string true "文章内容" minlength(3) maxlength(100000)
// @Param cover_image_url body string true "封面图片链接" minlength(3) maxlength(256)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app:BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	param.ID = convert.StrTo(c.Param("id")).MustUInt32()
	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签文章
// @Produce  json
// @Param id path int true "文章"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	response := app.NewResponse(c)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
