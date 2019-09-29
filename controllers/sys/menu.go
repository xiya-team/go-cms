package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type MenuController struct {
	controllers.BaseController
}

func (c *MenuController) Prepare() {

}

func (c *MenuController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewMenu().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "menu/index.html"
}

func (c *MenuController) Create() {
	if c.Ctx.Input.IsPost() {
		menuModel := models.NewMenu()
		//1.压入数据
		if err := c.ParseForm(menuModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(menuModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := menuModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.Menu{}
	c.TplName = c.ADMIN_TPL + "menu/create.html"
}

func (c *MenuController) Update() {
	id, _ := c.GetInt("id")
	menu, _ := models.NewMenu().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&menu); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(menu); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := menu.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = menu
	c.TplName = c.ADMIN_TPL + "menu/update.html"
}

func (c *MenuController) Delete() {
	menuModel := models.NewMenu()
	id, _ := c.GetInt("id")
	menuModel.Id = id
	if err := menuModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *MenuController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	menuModel := models.NewMenu()
	if err := menuModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

