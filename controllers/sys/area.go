package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type AreaController struct {
	controllers.BaseController
}

func (c *AreaController) Prepare() {

}

func (c *AreaController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")
		
		result, count := models.NewArea().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "area/index.html"
}

func (c *AreaController) Create() {
	if c.Ctx.Input.IsPost() {
		areaModel := models.NewArea()
		//1.压入数据
		if err := c.ParseForm(areaModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(areaModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := areaModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}
	
	c.Data["vo"] = models.Area{}
	c.TplName = c.ADMIN_TPL + "area/create.html"
}

func (c *AreaController) Update() {
	id, _ := c.GetInt("id")
	area, _ := models.NewArea().FindById(id)
	
	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&area); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(area); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := area.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}
	
	c.Data["vo"] = area
	c.TplName = c.ADMIN_TPL + "area/update.html"
}

func (c *AreaController) Delete() {
	areaModel := models.NewArea()
	id, _ := c.GetInt("id")
	areaModel.Id = id
	if err := areaModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *AreaController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}
	
	areaModel := models.NewArea()
	if err := areaModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

