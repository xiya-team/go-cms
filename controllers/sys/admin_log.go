package sys

import (
	"github.com/astaxie/beego/validation"
	"github.com/syyongx/php2go"
	"go-cms/controllers"
	"encoding/json"
	"go-cms/models"
    "go-cms/pkg/e"
	"log"
)

type AdminLogController struct {
	controllers.BaseController
}

func (c *AdminLogController) Prepare() {

}

/**
获取列表数据
 */
func (c *AdminLogController) Index() {
	if c.Ctx.Input.IsPost() {
		model := models.NewAdminLog()
		
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		err := json.Unmarshal(data, &model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		dataMap := make(map[string]interface{}, 0)

		//开始时间
		if !php2go.Empty(model.StartTime) {
			dataMap["start_time"] = model.StartTime
		}
		
		//结束时间
		if !php2go.Empty(model.EndTime) {
			dataMap["end_time"] = model.EndTime
		}
		
		var orderBy string = "created_at DESC"
		
		result, count,err := models.NewAdminLog().FindByMap((model.Page-1)*model.PageSize, model.PageSize, dataMap,orderBy)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}
		c.JsonResult(e.SUCCESS, "ok", result, count, model.Page, model.PageSize)
	}
}

/**
创建数据
*/
func (c *AdminLogController) Create() {
	if c.Ctx.Input.IsPost() {
		model := models.NewAdminLog()
        data := c.Ctx.Input.RequestBody
		//1.压入数据 json数据封装到对象中
		
		err := json.Unmarshal(data, model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
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
}

/**
更新数据
*/
func (c *AdminLogController) Update() {
	if c.Ctx.Input.IsPut() {
		model := models.NewAdminLog()
		data := c.Ctx.Input.RequestBody
		//json数据封装到对象中
		
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		post, err := models.NewAdminLog().FindById(model.Id)
		if err != nil||php2go.Empty(post) {
			c.JsonResult(e.ERROR, "没找到数据")
		}
		
		valid := validation.Validation{}
		if b, _ := valid.Valid(model); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(e.ERROR, "验证失败")
		}
		
		if _, err := model.Update(); err != nil {
			c.JsonResult(e.ERROR, "修改失败")
		}
		c.JsonResult(e.SUCCESS, "修改成功")
	}
}

/**
删除数据
*/
func (c *AdminLogController) Delete() {
    if c.Ctx.Input.IsDelete() {
		model := models.NewAdminLog()
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		post, err := models.NewAdminLog().FindById(model.Id)
		if err != nil||php2go.Empty(post) {
			c.JsonResult(e.ERROR, "没找到数据")
		}
		
		if err := model.Delete(); err != nil {
			c.JsonResult(e.ERROR, "删除失败")
		}
		c.JsonResult(e.SUCCESS, "删除成功")
	}
}

func (c *AdminLogController) BatchDelete() {
	model := models.NewAdminLog()

	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}

	if err := model.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

