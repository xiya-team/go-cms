package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type RoleMenuController struct {
	controllers.BaseController
}

func (c *RoleMenuController) Prepare() {

}

func (c *RoleMenuController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewRoleMenu().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "roleMenu/index.html"
}

func (c *RoleMenuController) Create() {
	if c.Ctx.Input.IsPost() {
		roleMenuModel := models.NewRoleMenu()
		//1.压入数据
		if err := c.ParseForm(roleMenuModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(roleMenuModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := roleMenuModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.RoleMenu{}
	c.TplName = c.ADMIN_TPL + "roleMenu/create.html"
}

func (c *RoleMenuController) Update() {
	id, _ := c.GetInt("id")
	roleMenu, _ := models.NewRoleMenu().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&roleMenu); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(roleMenu); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := roleMenu.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = roleMenu
	c.TplName = c.ADMIN_TPL + "roleMenu/update.html"
}

func (c *RoleMenuController) Delete() {
	roleMenuModel := models.NewRoleMenu()
	id, _ := c.GetInt("id")
	roleMenuModel.Id = id
	if err := roleMenuModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *RoleMenuController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	roleMenuModel := models.NewRoleMenu()
	if err := roleMenuModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

