package sys

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/syyongx/php2go"
	"go-cms/controllers"
	"go-cms/models"
	"go-cms/pkg/e"
	"log"
)

type PostController struct {
	controllers.BaseController
}

func (c *PostController) Prepare() {

}

func (c *PostController) Index() {
	if c.Ctx.Input.IsPost() {
		model := models.NewPost()
		
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		err := json.Unmarshal(data, model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		dataMap := make(map[string]interface{}, 0)
		
		if !php2go.Empty(model.PostCode) {
			dataMap["post_code"] = model.PostCode  //岗位编码
		}
		
		if !php2go.Empty(model.PostName) {
			dataMap["post_name"] = model.PostName  //岗位名称
		}
		if !php2go.Empty(model.Status) {
			dataMap["status"] = model.Status       //岗位状态
		}
		
		var orderBy string = "created_at DESC"
		
		result, count,err := model.FindByMap((model.Page-1)*model.PageSize, model.PageSize, dataMap,orderBy)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}
		c.JsonResult(e.SUCCESS, "ok", result, count)
	}
}

func (c *PostController) Create() {
	if c.Ctx.Input.IsPost() {
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		model := models.NewPost()
		err := json.Unmarshal(data, model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
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
}

func (c *PostController) Update() {
	if c.Ctx.Input.IsPut() {
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		model := models.NewPost()
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
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
		
		if _, err := model.Update(); err != nil {
			c.JsonResult(e.ERROR, "修改失败")
		}
		c.JsonResult(e.SUCCESS, "修改成功")
	}
}

func (c *PostController) Delete() {
	if c.Ctx.Input.IsDelete() {
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		model := models.NewPost()
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

func (c *PostController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}

	postModel := models.NewPost()
	if err := postModel.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

