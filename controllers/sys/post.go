package sys

import (
	"github.com/syyongx/php2go"
	"go-cms/controllers"
	"go-cms/models"
	"github.com/astaxie/beego/validation"
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
		page, _ := c.GetInt("page",1)
		limit, _ := c.GetInt("limit",10)
		
		post := models.NewPost()
		
		dataMap := make(map[string]interface{}, 0)
		
		if !php2go.Empty(post.PostCode) {
			dataMap["post_code"] = post.PostCode  //岗位编码
		}
		
		if !php2go.Empty(post.PostName) {
			dataMap["post_name"] = post.PostName  //岗位名称
		}
		if !php2go.Empty(post.Status) {
			dataMap["status"] = post.Status       //岗位状态
		}
		
		var orderBy string = "created_at DESC"
		
		result, count,err := models.NewPost().FindByMap((page-1)*limit, limit, dataMap,orderBy)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}
		c.JsonResult(e.SUCCESS, "ok", result, count)
		
	}
}

func (c *PostController) Create() {
	if c.Ctx.Input.IsPost() {
		postModel := models.NewPost()
		//1.压入数据
		if err := c.ParseForm(postModel); err != nil {
			c.JsonResult(e.ERROR, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(postModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(e.ERROR, "验证失败")
		}
		//3.插入数据
		if _, err := postModel.Create(); err != nil {
			c.JsonResult(e.ERROR, "创建失败")
		}
		c.JsonResult(e.SUCCESS, "添加成功")
	}
}

func (c *PostController) Update() {
	id, _ := c.GetInt("id")
	post, _ := models.NewPost().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&post); err != nil {
			c.JsonResult(e.ERROR, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(post); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(e.ERROR, "验证失败")
		}
		//3
		if _, err := post.Update(); err != nil {
			c.JsonResult(e.ERROR, "修改失败")
		}
		c.JsonResult(e.SUCCESS, "修改成功")
	}
}

func (c *PostController) Delete() {
	postModel := models.NewPost()
	id, _ := c.GetInt("id")
	postModel.Id = id
	if err := postModel.Delete(); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
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

