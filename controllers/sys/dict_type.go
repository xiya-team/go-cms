package sys

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/xiya-team/helpers"
	"go-cms/common"
	"go-cms/controllers"
	"go-cms/models"
	"go-cms/pkg/e"
	"go-cms/pkg/util"
	"go-cms/validations"
	"log"
	"reflect"
	"strings"
)

type DictTypeController struct {
	controllers.BaseController
}

func (c *DictTypeController) Prepare() {

}

/**
获取列表数据
 */
func (c *DictTypeController) Index() {
	if c.Ctx.Input.IsPost() {
		model := models.NewDictType()
		
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		err := json.Unmarshal(data, &model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		dataMap := make(map[string]interface{}, 0)
		
		if !helpers.Empty(model.DictName) {
			dataMap["dict_name"] = model.DictName
		}
		
		if !helpers.Empty(model.DictType) {
			dataMap["dict_type"] = model.DictType
		}

		//开始时间
		if !helpers.Empty(model.StartTime) {
			dataMap["start_time"] = model.StartTime
		}
		
		//结束时间
		if !helpers.Empty(model.EndTime) {
			dataMap["end_time"] = model.EndTime
		}
		
		//状态
		if !helpers.Empty(model.Status) {
			dataMap["status"] = model.Status
		}

		//查询字段
		if !helpers.Empty(model.Fields) {
			dataMap["fields"] = model.Fields
		}

		if helpers.Empty(model.Page) {
			model.Page = 1
		}else{
			if model.Page <= 0 {
				model.Page = 1
			}
		}

		if helpers.Empty(model.PageSize) {
			model.PageSize = 10
		}else {
			if model.Page <= 0 {
				model.Page = 10
			}
		}

		var orderBy string
		if !helpers.Empty(model.OrderColumnName) && !helpers.Empty(model.OrderType){
			orderBy = strings.Join([]string{model.OrderColumnName,model.OrderType}," ")
		}else {
			orderBy = "created_at DESC"
		}
		
		result, count,err := models.NewDictType().FindByMap((model.Page-1)*model.PageSize, model.PageSize, dataMap,orderBy)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}

		if !helpers.Empty(model.Fields){
			fields := strings.Split(model.Fields, ",")
			lists := c.FormatData(fields,result)
			c.JsonResult(e.SUCCESS, "ok", lists, count, model.Page, model.PageSize)
		}else {
			c.JsonResult(e.SUCCESS, "ok", result, count, model.Page, model.PageSize)
		}
	}
}

/**
创建数据
*/
func (c *DictTypeController) Create() {
	if c.Ctx.Input.IsPut() {
		model := models.NewDictType()
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
		if !helpers.Empty(message){
			c.JsonResult(e.ERROR, message)
		}

		user,_ := model.FindByDictType(model.DictType)
		if !helpers.Empty(user) {
			c.JsonResult(e.ERROR, "字典类型已存在!")
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
		model.CreateBy = common.UserId
		if _, err := model.Create(); err != nil {
			c.JsonResult(e.ERROR, "字典类型已存在！")
		}
		c.JsonResult(e.SUCCESS, "添加成功！")
	}
}

/**
更新数据
*/
func (c *DictTypeController) Update() {
	model := models.NewDictType()
	data := c.Ctx.Input.RequestBody
	//json数据封装到对象中

	err := json.Unmarshal(data, model)

	if err != nil {
		c.JsonResult(e.ERROR, err.Error())
	}

	if c.Ctx.Input.IsPut() {
		post, err := models.NewDictType().FindById(model.Id)
		if err != nil||helpers.Empty(post) {
			c.JsonResult(e.ERROR, "没找到数据")
		}

		//2.验证
		UserValidations := validations.BaseValidations{}
		message := UserValidations.Check(model)
		if !helpers.Empty(message){
			c.JsonResult(e.ERROR, message)
		}

		dict_type,_ := model.FindByDictType(model.DictType)
		if !helpers.Empty(dict_type) && dict_type.Id!=model.Id {
			c.JsonResult(e.ERROR, "字典类型已存在!")
		}

		model.UpdateBy = common.UserId
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
func (c *DictTypeController) Delete() {
    if c.Ctx.Input.IsDelete() {
		model := models.NewDictType()
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		post, err := models.NewDictType().FindById(model.Id)
		if err != nil||helpers.Empty(post) {
			c.JsonResult(e.ERROR, "没找到数据")
		}
		
		if err := model.Delete(); err != nil {
			c.JsonResult(e.ERROR, "删除失败")
		}
		c.JsonResult(e.SUCCESS, "删除成功")
	}
}

func (c *DictTypeController) BatchDelete() {
	model := models.NewDictType()

	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}

	if err := model.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

func (c *DictTypeController) FormatData(fields []string,result []models.DictType) (res interface{}) {
	lists := make([]map[string]interface{},0)

	for key,item:=range fields {
		fields[key] = util.ToFirstWordsUp(item)
	}

	for _, value := range result {
		tmp := make(map[string]interface{}, 0)
		t := reflect.TypeOf(value)
		v := reflect.ValueOf(value)
		for k := 0; k < t.NumField(); k++ {
			if helpers.InArray(t.Field(k).Name,fields){
				tmp[util.ToFirstWordsDown(t.Field(k).Name)] = v.Field(k).Interface()
			}
		}
		lists = append(lists,tmp)
	}
	return lists
}

func (c *DictTypeController) FindById() {
	if c.Ctx.Input.IsPost() {
		model := models.NewDictType()

		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		err := json.Unmarshal(data, &model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}

		dataMap := make(map[string]interface{}, 0)

		if !helpers.Empty(model.Id) {
			dataMap["id"] = model.Id
		}

		//查询字段
		if !helpers.Empty(model.Fields) {
			dataMap["fields"] = model.Fields
		}

		result, err := models.NewDictType().FindById(model.Id)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}

		c.JsonResult(e.SUCCESS, "ok", result)
	}
}