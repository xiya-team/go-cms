package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type ConfigssController struct {
	controllers.BaseController
}

func (c *ConfigssController) Prepare() {

}

func (c *ConfigssController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewConfigss().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "configss/index.html"
}

func (c *ConfigssController) Create() {
	if c.Ctx.Input.IsPost() {
		configssModel := models.NewConfigss()
		//1.压入数据
		if err := c.ParseForm(configssModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(configssModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := configssModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.Configss{}
	c.TplName = c.ADMIN_TPL + "configss/create.html"
}

func (c *ConfigssController) Update() {
	id, _ := c.GetInt("id")
	configss, _ := models.NewConfigss().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&configss); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(configss); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := configss.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = configss
	c.TplName = c.ADMIN_TPL + "configss/update.html"
}

func (c *ConfigssController) Delete() {
	configssModel := models.NewConfigss()
	id, _ := c.GetInt("id")
	configssModel.Id = id
	if err := configssModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *ConfigssController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	configssModel := models.NewConfigss()
	if err := configssModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

