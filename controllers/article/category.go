package article

import (
	"github.com/astaxie/beego/validation"
	"go-cms/controllers"
	"go-cms/models"
	"log"
)

type CategoryController struct {
	controllers.BaseController
}

func (c *CategoryController) Prepare() {

}

func (c *CategoryController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewCategory().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "category/index.html"
}

func (c *CategoryController) Create() {
	if c.Ctx.Input.IsPost() {
		categoryModel := models.NewCategory()
		//1.压入数据
		if err := c.ParseForm(categoryModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(categoryModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := categoryModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.Category{}
	c.TplName = c.ADMIN_TPL + "category/create.html"
}

func (c *CategoryController) Update() {
	id, _ := c.GetInt("id")
	category, _ := models.NewCategory().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&category); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(category); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := category.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = category
	c.TplName = c.ADMIN_TPL + "category/update.html"
}

func (c *CategoryController) Delete() {
	categoryModel := models.NewCategory()
	id, _ := c.GetInt("id")
	categoryModel.Id = id
	if err := categoryModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *CategoryController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	categoryModel := models.NewCategory()
	if err := categoryModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

