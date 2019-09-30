package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type RoleController struct {
	controllers.BaseController
}

func (c *RoleController) Prepare() {

}

func (c *RoleController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewRole().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "role/index.html"
}

func (c *RoleController) Create() {
	if c.Ctx.Input.IsPost() {
		roleModel := models.NewRole()
		//1.压入数据
		if err := c.ParseForm(roleModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(roleModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := roleModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.Role{}
	c.TplName = c.ADMIN_TPL + "role/create.html"
}

func (c *RoleController) Update() {
	id, _ := c.GetInt("id")
	role, _ := models.NewRole().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&role); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(role); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := role.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = role
	c.TplName = c.ADMIN_TPL + "role/update.html"
}

func (c *RoleController) Delete() {
	roleModel := models.NewRole()
	id, _ := c.GetInt("id")
	roleModel.Id = id
	if err := roleModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *RoleController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	roleModel := models.NewRole()
	if err := roleModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

