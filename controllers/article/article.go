package article

import (
	"github.com/astaxie/beego/validation"
	"go-cms/controllers"
	"go-cms/models"
	"log"
)

type ArticleController struct {
	controllers.BaseController
}

func (c *ArticleController) Prepare() {

}

func (c *ArticleController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewArticle().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "article/index.html"
}

func (c *ArticleController) Create() {
	if c.Ctx.Input.IsPost() {
		articleModel := models.NewArticle()
		//1.压入数据
		if err := c.ParseForm(articleModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(articleModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := articleModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.Article{}
	c.TplName = c.ADMIN_TPL + "article/create.html"
}

func (c *ArticleController) Update() {
	id, _ := c.GetInt("id")
	article, _ := models.NewArticle().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&article); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(article); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := article.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = article
	c.TplName = c.ADMIN_TPL + "article/update.html"
}

func (c *ArticleController) Delete() {
	articleModel := models.NewArticle()
	id, _ := c.GetInt("id")
	articleModel.Id = id
	if err := articleModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *ArticleController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	articleModel := models.NewArticle()
	if err := articleModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

