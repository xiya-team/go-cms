package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type RoleDeptController struct {
	controllers.BaseController
}

func (c *RoleDeptController) Prepare() {

}

func (c *RoleDeptController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewRoleDept().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "roleDept/index.html"
}

func (c *RoleDeptController) Create() {
	if c.Ctx.Input.IsPost() {
		roleDeptModel := models.NewRoleDept()
		//1.压入数据
		if err := c.ParseForm(roleDeptModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(roleDeptModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := roleDeptModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.RoleDept{}
	c.TplName = c.ADMIN_TPL + "roleDept/create.html"
}

func (c *RoleDeptController) Update() {
	id, _ := c.GetInt("id")
	roleDept, _ := models.NewRoleDept().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&roleDept); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(roleDept); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := roleDept.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = roleDept
	c.TplName = c.ADMIN_TPL + "roleDept/update.html"
}

func (c *RoleDeptController) Delete() {
	roleDeptModel := models.NewRoleDept()
	id, _ := c.GetInt("id")
	roleDeptModel.Id = id
	if err := roleDeptModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *RoleDeptController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	roleDeptModel := models.NewRoleDept()
	if err := roleDeptModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

