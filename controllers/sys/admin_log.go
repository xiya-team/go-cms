package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type AdminLogController struct {
	controllers.BaseController
}

func (c *AdminLogController) Prepare() {

}

func (c *AdminLogController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewAdminLog().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "adminLog/index.html"
}

func (c *AdminLogController) Create() {
	if c.Ctx.Input.IsPost() {
		adminLogModel := models.NewAdminLog()
		//1.压入数据
		if err := c.ParseForm(adminLogModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(adminLogModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := adminLogModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.AdminLog{}
	c.TplName = c.ADMIN_TPL + "adminLog/create.html"
}

func (c *AdminLogController) Update() {
	id, _ := c.GetInt("id")
	adminLog, _ := models.NewAdminLog().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&adminLog); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(adminLog); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := adminLog.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = adminLog
	c.TplName = c.ADMIN_TPL + "adminLog/update.html"
}

func (c *AdminLogController) Delete() {
	adminLogModel := models.NewAdminLog()
	id, _ := c.GetInt("id")
	adminLogModel.Id = id
	if err := adminLogModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *AdminLogController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	adminLogModel := models.NewAdminLog()
	if err := adminLogModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

