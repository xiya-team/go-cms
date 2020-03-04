package generate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var controllerTpl = `package $path$

import (
	"github.com/astaxie/beego/validation"
	"github.com/xiya-team/helpers"
	"go-cms/controllers"
	"encoding/json"
	"go-cms/models"
    "go-cms/common"
    "go-cms/validations"
    "go-cms/pkg/e"
	"log"
	"strings"
)

type CategoryController struct {
	controllers.BaseController
}

func (c *CategoryController) Prepare() {

}

/**
获取列表数据
 */
func (c *CategoryController) Index() {
	if c.Ctx.Input.IsPost() {
		model := models.NewCategory()
		
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
		
		//状态
		if !php2go.Empty(model.Status) {
			dataMap["status"] = model.Status
		}
		
		if php2go.Empty(model.Page) {
			model.Page = 1
		}else{
			if model.Page <= 0 {
				model.Page = 1
			}
		}

		if php2go.Empty(model.PageSize) {
			model.PageSize = 10
		}else {
			if model.Page <= 0 {
				model.Page = 10
			}
		}

		var orderBy string
		if !php2go.Empty(model.OrderColumnName) && !php2go.Empty(model.OrderType){
			orderBy = strings.Join([]string{model.OrderColumnName,model.OrderType}," ")
		}else {
			orderBy = "created_at DESC"
		}
		
		result, count,err := model.FindByMap((model.Page-1)*model.PageSize, model.PageSize, dataMap,orderBy)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}
		c.JsonResult(e.SUCCESS, "ok", result, count, model.Page, model.PageSize)
	}
}

/**
创建数据
*/
func (c *CategoryController) Create() {
	if c.Ctx.Input.IsPost() {
		model := models.NewCategory()
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
		UserValidations := validations.BaseValidations{}
		message := UserValidations.Check(model)
		if !php2go.Empty(message){
			c.JsonResult(e.ERROR, message)
		}

		//3.插入数据
		model.CreatedBy = common.UserId
		if _, err := model.Create(); err != nil {
			c.JsonResult(e.ERROR, "创建失败")
		}
		c.JsonResult(e.SUCCESS, "添加成功")
	}
}

/**
更新数据
*/
func (c *CategoryController) Update() {
    model := models.NewCategory()
	data := c.Ctx.Input.RequestBody
	//json数据封装到对象中
	
	err := json.Unmarshal(data, model)
	
	if err != nil {
		c.JsonResult(e.ERROR, err.Error())
	}

	if c.Ctx.Input.IsPut() {
		post, err := model.FindById(model.Id)
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
		
		model.UpdatedBy = common.UserId
		if _, err := model.Update(); err != nil {
			c.JsonResult(e.ERROR, "修改失败")
		}
		c.JsonResult(e.SUCCESS, "修改成功")
	}

	//get
	if c.Ctx.Input.IsPost() {
		res,err := model.FindById(model.Id)
		if err != nil{
			c.JsonResult(e.ERROR, "获取失败")
		}
		c.JsonResult(e.SUCCESS, "获取成功",res)
	}
}

/**
删除数据
*/
func (c *CategoryController) Delete() {
    if c.Ctx.Input.IsDelete() {
		model := models.NewCategory()
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		post, err := model.FindById(model.Id)
		if err != nil||php2go.Empty(post) {
			c.JsonResult(e.ERROR, "没找到数据")
		}
		
		if err := model.Delete(); err != nil {
			c.JsonResult(e.ERROR, "删除失败")
		}
		c.JsonResult(e.SUCCESS, "删除成功")
	}
}

func (c *CategoryController) BatchDelete() {
	model := models.NewCategory()

	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}

	if err := model.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

`

func CreateController(controllerPath, tableName string) {
	controllerData := ReplaceStr(tableName, controllerPath,controllerTpl,"")

	if err := os.MkdirAll(path.Clean(controllerPath), 777); err != nil {
		log.Println("控制器文件创建失败")
	}

	if err := ioutil.WriteFile(path.Join(controllerPath, fmt.Sprintf("%s.go", tableName)), []byte(controllerData), os.ModeType); err != nil {
		log.Println(err)
	}
}