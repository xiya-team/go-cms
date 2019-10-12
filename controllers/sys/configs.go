package sys

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/syyongx/php2go"
	"go-cms/controllers"
	"go-cms/models"
	"go-cms/pkg/e"
	"log"
)

type ConfigsController struct {
	controllers.BaseController
}

func (c *ConfigsController) Prepare() {

}

func (c *ConfigsController) Index() {
	if c.Ctx.Input.IsPost() {
		page, _ := c.GetInt("page",1)
		limit, _ := c.GetInt("limit",10)
		
		configs := models.NewConfigs()
		
		dataMap := make(map[string]interface{}, 0)
		
		if !php2go.Empty(configs.ConfigType) {
			dataMap["config_type"] = configs.ConfigType
		}
		
		if !php2go.Empty(configs.StartTime) {
			dataMap["start_time"] = configs.StartTime
		}
		if !php2go.Empty(configs.EndTime) {
			dataMap["end_time"] = configs.EndTime
		}
		
		var orderBy string = "created_at DESC"
		
		result, count,err := models.NewConfigs().FindByMap((page-1)*limit, limit, dataMap,orderBy)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}
		c.JsonResult(e.SUCCESS, "ok", result, count)
	}
}

func (c *ConfigsController) Create() {
	if c.Ctx.Input.IsPost() {
		configsModel := models.NewConfigs()
		//1.压入数据
		if err := c.ParseForm(configsModel); err != nil {
			c.JsonResult(e.ERROR, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(configsModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(e.ERROR, "验证失败")
		}
		//3.插入数据
		if _, err := configsModel.Create(); err != nil {
			c.JsonResult(e.ERROR, "创建失败")
		}
		c.JsonResult(e.SUCCESS, "添加成功")
	}
}

func (c *ConfigsController) Update() {
	
	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("id")
		configs, _ := models.NewConfigs().FindById(id)
		//1
		if err := c.ParseForm(&configs); err != nil {
			c.JsonResult(e.ERROR, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(configs); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(e.ERROR, "验证失败")
		}
		//3
		if _, err := configs.Update(); err != nil {
			logs.Debug(err.Error())
			c.JsonResult(e.ERROR, "修改失败")
		}
		c.JsonResult(e.SUCCESS, "修改成功")
	}
}

func (c *ConfigsController) Delete() {
	configsModel := models.NewConfigs()
	id, _ := c.GetInt("id")
	configsModel.Id = id
	if err := configsModel.Delete(); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

func (c *ConfigsController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}
	
	configsModel := models.NewConfigs()
	if err := configsModel.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}