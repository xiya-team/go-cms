package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type UserPostController struct {
	controllers.BaseController
}

func (c *UserPostController) Prepare() {

}

func (c *UserPostController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewUserPost().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "userPost/index.html"
}

func (c *UserPostController) Create() {
	if c.Ctx.Input.IsPost() {
		userPostModel := models.NewUserPost()
		//1.压入数据
		if err := c.ParseForm(userPostModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(userPostModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := userPostModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.UserPost{}
	c.TplName = c.ADMIN_TPL + "userPost/create.html"
}

func (c *UserPostController) Update() {
	id, _ := c.GetInt("id")
	userPost, _ := models.NewUserPost().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&userPost); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(userPost); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := userPost.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = userPost
	c.TplName = c.ADMIN_TPL + "userPost/update.html"
}

func (c *UserPostController) Delete() {
	userPostModel := models.NewUserPost()
	id, _ := c.GetInt("id")
	userPostModel.Id = id
	if err := userPostModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *UserPostController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	userPostModel := models.NewUserPost()
	if err := userPostModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

