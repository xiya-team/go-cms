package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type DictTypeController struct {
	controllers.BaseController
}

func (c *DictTypeController) Prepare() {

}

func (c *DictTypeController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewDictType().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "dictType/index.html"
}

func (c *DictTypeController) Create() {
	if c.Ctx.Input.IsPost() {
		dictTypeModel := models.NewDictType()
		//1.压入数据
		if err := c.ParseForm(dictTypeModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(dictTypeModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := dictTypeModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.DictType{}
	c.TplName = c.ADMIN_TPL + "dictType/create.html"
}

func (c *DictTypeController) Update() {
	id, _ := c.GetInt("id")
	dictType, _ := models.NewDictType().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&dictType); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(dictType); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := dictType.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = dictType
	c.TplName = c.ADMIN_TPL + "dictType/update.html"
}

func (c *DictTypeController) Delete() {
	dictTypeModel := models.NewDictType()
	id, _ := c.GetInt("id")
	dictTypeModel.Id = id
	if err := dictTypeModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *DictTypeController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	dictTypeModel := models.NewDictType()
	if err := dictTypeModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

