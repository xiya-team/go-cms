package sys

import (
	"encoding/json"
	"github.com/xiya-team/helpers"
	"go-cms/common"
	"go-cms/controllers"
	"go-cms/models"
	"go-cms/pkg/e"
	"go-cms/pkg/util"
	"go-cms/validations"
	"reflect"
	"strings"
)

type DictDataController struct {
	controllers.BaseController
}

func (c *DictDataController) Prepare() {

}

/**
获取列表数据
 */
func (c *DictDataController) Index() {
	if c.Ctx.Input.IsPost() {
		model := models.NewDictData()
		
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		err := json.Unmarshal(data, &model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		dataMap := make(map[string]interface{}, 0)
		if !helpers.Empty(model.DictId) {
			dataMap["dict_id"] = model.DictId
		}
		
		if !helpers.Empty(model.DictLabel) {
			dataMap["dict_label"] = model.DictLabel
		}

		if !helpers.Empty(model.DictType) {
			dataMap["dict_type"] = model.DictType
		}

		if !helpers.Empty(model.DictValueType) {
			dataMap["dict_value_type"] = model.DictValueType
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
		
		result, count,err := models.NewDictData().FindByMap((model.Page-1)*model.PageSize, model.PageSize, dataMap,orderBy)

		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}

		maps := make(map[string]interface{})
		maps["page"] = util.Pages(count, model.Page, model.PageSize)

		if !helpers.Empty(model.DictId){
			dict_type_models := models.NewDictType()
			dict_type,_ := dict_type_models.FindById(model.DictId)
			maps["dict_value_type"] = dict_type.DictValueType
		}
		if !helpers.Empty(model.Fields){
			fields := strings.Split(model.Fields, ",")
			lists := c.FormatData(fields,result)
			maps["list"] = lists
			c.JsonResult(e.SUCCESS, "ok", maps)
		}else {
			maps["list"] = result
			c.JsonResult(e.SUCCESS, "ok", maps)
		}
	}
}

/**
创建数据
*/
func (c *DictDataController) Create() {
	if c.Ctx.Input.IsPut() {
		model := models.NewDictData()
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

		whereMap := make(map[string]interface{}, 0)
		whereMap["dict_id"] = model.DictId

		dictTypeModel := models.NewDictType()
		dictType,_ := dictTypeModel.FindById(model.DictId)

		//数值类型 1 数值 2 字符串
		if dictType.DictValueType == 1{
			whereMap["dict_number"] = model.DictNumber
		}else {
			whereMap["dict_value"] = model.DictValue
		}

		model.DictValueType = dictType.DictValueType
		model.DictType = dictType.DictType

		//3.插入数据
		model.CreateBy = common.UserId
		if _, err := model.Create(); err != nil {
			c.JsonResult(e.ERROR, "创建失败")
		}
		c.JsonResult(e.SUCCESS, "添加成功")
	}
}

/**
更新数据
*/
func (c *DictDataController) Update() {
	model := models.NewDictData()
	data := c.Ctx.Input.RequestBody
	//json数据封装到对象中

	err := json.Unmarshal(data, model)

	if err != nil {
		c.JsonResult(e.ERROR, err.Error())
	}

	if c.Ctx.Input.IsPut() {
		dict_data, err := models.NewDictData().FindById(model.Id)
		if err != nil||helpers.Empty(dict_data) {
			c.JsonResult(e.ERROR, "没找到数据")
		}

		//2.验证
		UserValidations := validations.BaseValidations{}
		message := UserValidations.Check(model)
		if !helpers.Empty(message){
			c.JsonResult(e.ERROR, message)
		}

		whereMap := make(map[string]interface{}, 0)
		whereMap["dict_id"] = model.DictId

		dictTypeModel := models.NewDictType()
		dictType,_ := dictTypeModel.FindById(model.DictId)

		//数值类型 1 数值 2 字符串
		if dictType.DictValueType == 1{
			whereMap["dict_number"] = model.DictNumber
		}else {
			whereMap["dict_value"] = model.DictValue
		}

		model.DictValueType = dictType.DictValueType
		model.DictType = dictType.DictType

		tmp,_:= model.FindWhere(whereMap)
		if !helpers.Empty(tmp) && tmp.Id!=dict_data.Id {
			c.JsonResult(e.ERROR, "数据重复！")
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
func (c *DictDataController) Delete() {
    if c.Ctx.Input.IsDelete() {
		model := models.NewDictData()
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		post, err := models.NewDictData().FindById(model.Id)
		if err != nil||helpers.Empty(post) {
			c.JsonResult(e.ERROR, "没找到数据")
		}
		
		if err := model.Delete(); err != nil {
			c.JsonResult(e.ERROR, "删除失败")
		}
		c.JsonResult(e.SUCCESS, "删除成功")
	}
}

func (c *DictDataController) BatchDelete() {
	model := models.NewDictData()

	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}

	if err := model.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

func (c *DictDataController) FormatData(fields []string,result []models.DictData) (res interface{}) {
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