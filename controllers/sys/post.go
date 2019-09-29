package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type PostController struct {
	controllers.BaseController
}

func (c *PostController) Prepare() {

}

func (c *PostController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewPost().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "post/index.html"
}

func (c *PostController) Create() {
	if c.Ctx.Input.IsPost() {
		postModel := models.NewPost()
		//1.压入数据
		if err := c.ParseForm(postModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(postModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := postModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.Post{}
	c.TplName = c.ADMIN_TPL + "post/create.html"
}

func (c *PostController) Update() {
	id, _ := c.GetInt("id")
	post, _ := models.NewPost().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&post); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(post); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := post.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = post
	c.TplName = c.ADMIN_TPL + "post/update.html"
}

func (c *PostController) Delete() {
	postModel := models.NewPost()
	id, _ := c.GetInt("id")
	postModel.Id = id
	if err := postModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *PostController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	postModel := models.NewPost()
	if err := postModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

