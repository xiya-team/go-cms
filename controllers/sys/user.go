package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Prepare() {

}

func (c *UserController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewUser().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "user/index.html"
}

func (c *UserController) Create() {
	if c.Ctx.Input.IsPost() {
		userModel := models.NewUser()
		//1.压入数据
		if err := c.ParseForm(userModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(userModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := userModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.User{}
	c.TplName = c.ADMIN_TPL + "user/create.html"
}

func (c *UserController) Update() {
	id, _ := c.GetInt("id")
	user, _ := models.NewUser().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&user); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(user); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := user.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = user
	c.TplName = c.ADMIN_TPL + "user/update.html"
}

func (c *UserController) Delete() {
	userModel := models.NewUser()
	id, _ := c.GetInt("id")
	userModel.Id = id
	if err := userModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *UserController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	userModel := models.NewUser()
	if err := userModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

