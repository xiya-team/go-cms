package sys

import (
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type DictDataController struct {
	controllers.BaseController
}

func (c *DictDataController) Prepare() {

}

func (c *DictDataController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewDictData().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "dictData/index.html"
}

func (c *DictDataController) Create() {
	if c.Ctx.Input.IsPost() {
		dictDataModel := models.NewDictData()
		//1.压入数据
		if err := c.ParseForm(dictDataModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(dictDataModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := dictDataModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.DictData{}
	c.TplName = c.ADMIN_TPL + "dictData/create.html"
}

func (c *DictDataController) Update() {
	id, _ := c.GetInt("id")
	dictData, _ := models.NewDictData().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&dictData); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(dictData); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := dictData.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = dictData
	c.TplName = c.ADMIN_TPL + "dictData/update.html"
}

func (c *DictDataController) Delete() {
	dictDataModel := models.NewDictData()
	id, _ := c.GetInt("id")
	dictDataModel.Id = id
	if err := dictDataModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *DictDataController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	dictDataModel := models.NewDictData()
	if err := dictDataModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

