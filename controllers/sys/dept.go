package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type DeptController struct {
	controllers.BaseController
}

func (c *DeptController) Prepare() {

}

func (c *DeptController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewDept().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "dept/index.html"
}

func (c *DeptController) Create() {
	if c.Ctx.Input.IsPost() {
		deptModel := models.NewDept()
		//1.压入数据
		if err := c.ParseForm(deptModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(deptModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := deptModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.Dept{}
	c.TplName = c.ADMIN_TPL + "dept/create.html"
}

func (c *DeptController) Update() {
	id, _ := c.GetInt("id")
	dept, _ := models.NewDept().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&dept); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(dept); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := dept.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = dept
	c.TplName = c.ADMIN_TPL + "dept/update.html"
}

func (c *DeptController) Delete() {
	deptModel := models.NewDept()
	id, _ := c.GetInt("id")
	deptModel.Id = id
	if err := deptModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *DeptController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	deptModel := models.NewDept()
	if err := deptModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

