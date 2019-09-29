package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type UserRoleController struct {
	controllers.BaseController
}

func (c *UserRoleController) Prepare() {

}

func (c *UserRoleController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewUserRole().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "userRole/index.html"
}

func (c *UserRoleController) Create() {
	if c.Ctx.Input.IsPost() {
		userRoleModel := models.NewUserRole()
		//1.压入数据
		if err := c.ParseForm(userRoleModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(userRoleModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := userRoleModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.UserRole{}
	c.TplName = c.ADMIN_TPL + "userRole/create.html"
}

func (c *UserRoleController) Update() {
	id, _ := c.GetInt("id")
	userRole, _ := models.NewUserRole().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&userRole); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(userRole); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := userRole.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = userRole
	c.TplName = c.ADMIN_TPL + "userRole/update.html"
}

func (c *UserRoleController) Delete() {
	userRoleModel := models.NewUserRole()
	id, _ := c.GetInt("id")
	userRoleModel.Id = id
	if err := userRoleModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *UserRoleController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	userRoleModel := models.NewUserRole()
	if err := userRoleModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

