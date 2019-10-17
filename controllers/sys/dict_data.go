package sys

import (
	"github.com/syyongx/php2go"
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
	"go-cms/pkg/e"
	"log"
)

type DictDataController struct {
	controllers.BaseController
}

func (c *DictDataController) Prepare() {

}

func (c *DictDataController) Index() {
    if c.Ctx.Input.IsPost() {
		model := models.NewDictData()
		
		dataMap := make(map[string]interface{}, 0)
		
		if !php2go.Empty(model.DictId) {
			dataMap["dict_id"] = model.DictId
		}
	
	    if !php2go.Empty(model.DictLabel) {
		    dataMap["dict_label"] = model.DictLabel
	    }
		
	    if !php2go.Empty(model.Status) {
		    dataMap["status"] = model.Status
	    }
		
		if !php2go.Empty(model.StartTime) {
			dataMap["start_time"] = model.StartTime
		}
		if !php2go.Empty(model.EndTime) {
			dataMap["end_time"] = model.EndTime
		}
		
		var orderBy string = "created_at DESC"
	
	    result, count,err := model.FindByMap((model.Page-1)*model.PageSize, model.PageSize, dataMap,orderBy)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}
		c.JsonResult(e.SUCCESS, "ok", result, count)
	}
}

func (c *DictDataController) Create() {
	if c.Ctx.Input.IsPost() {
		model := models.NewDictData()
		//1.压入数据
		if err := c.ParseForm(model); err != nil {
			c.JsonResult(e.ERROR, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(model); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(e.ERROR, "验证失败")
		}
		//3.插入数据
		if _, err := model.Create(); err != nil {
			c.JsonResult(e.ERROR, "创建失败")
		}
		c.JsonResult(e.SUCCESS, "添加成功")
	}
	c.Data["vo"] = models.DictData{}
	c.TplName = c.ADMIN_TPL + "dictData/create.html"
}

func (c *DictDataController) Update() {
	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("id")
		model, _ := models.NewDictData().FindById(id)
		//1
		if err := c.ParseForm(&model); err != nil {
			c.JsonResult(e.ERROR, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(model); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(e.ERROR, "验证失败")
		}
		//3
		if _, err := model.Update(); err != nil {
			c.JsonResult(e.ERROR, "修改失败")
		}
		c.JsonResult(e.SUCCESS, "修改成功")
	}
}

func (c *DictDataController) Delete() {
    model := models.NewDictData()
	id, _ := c.GetInt("id")
	model.Id = id
	if err := model.Delete(); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

func (c *DictDataController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}
	
	model := models.NewDictData()
	if err := model.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

