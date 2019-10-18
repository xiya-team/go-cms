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

type MenuController struct {
	controllers.BaseController
}

func (c *MenuController) Prepare() {

}

/**
获取列表数据
 */
func (c *MenuController) Index() {
	if c.Ctx.Input.IsPost() {
		model := models.NewMenu()
		
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		err := json.Unmarshal(data, &model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		dataMap := make(map[string]interface{}, 0)
		
		if !php2go.Empty(model.Visible) {
			dataMap["visible"] = model.Visible
		}
		
		if !php2go.Empty(model.MenuName) {
			dataMap["menu_name"] = model.MenuName
		}

		//开始时间
		if !php2go.Empty(model.StartTime) {
			dataMap["start_time"] = model.StartTime
		}
		
		//结束时间
		if !php2go.Empty(model.EndTime) {
			dataMap["end_time"] = model.EndTime
		}
		
		var orderBy string = "created_at DESC"
		
		result, count,err := models.NewMenu().FindByMap((model.Page-1)*model.PageSize, model.PageSize, dataMap,orderBy)
		if err != nil{
			c.JsonResult(e.ERROR, "获取数据失败")
		}
		c.JsonResult(e.SUCCESS, "ok", result, count, model.Page, model.PageSize)
	}
}

/**
创建数据
*/
func (c *MenuController) Create() {
	if c.Ctx.Input.IsPost() {
		model := models.NewMenu()
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
func (c *MenuController) Update() {
	if c.Ctx.Input.IsPut() {
		model := models.NewMenu()
		data := c.Ctx.Input.RequestBody
		//json数据封装到对象中
		
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		post, err := models.NewMenu().FindById(model.Id)
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
func (c *MenuController) Delete() {
    if c.Ctx.Input.IsDelete() {
		model := models.NewMenu()
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		post, err := models.NewMenu().FindById(model.Id)
		if err != nil||php2go.Empty(post) {
			c.JsonResult(e.ERROR, "没找到数据")
		}
		
		if err := model.Delete(); err != nil {
			c.JsonResult(e.ERROR, "删除失败")
		}
		c.JsonResult(e.SUCCESS, "删除成功")
	}
}

func (c *MenuController) BatchDelete() {
	model := models.NewMenu()

	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}

	if err := model.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

func (c *MenuController) Menus()  {
	model := models.NewMenu()
	data := c.Ctx.Input.RequestBody
	//json数据封装到user对象中
	
	err := json.Unmarshal(data, model)
	
	if err != nil {
		c.JsonResult(e.ERROR, err.Error())
	}
	
	
	if php2go.Empty(model.ParentId){
		menuData,_ := model.FindAll()
		c.JsonResult(e.SUCCESS, "删除成功",constructMenuTrees(menuData,0))
	}else {
		menuData,_ := model.FindAllByParentId()
		c.JsonResult(e.SUCCESS, "删除成功",constructMenuTrees(menuData,0))
	}
	
	
	
	

}


type Item struct {
	MenuName     string `json:"menu_name"`
	ID           int    `json:"id"`
	Url          string `json:"url"`
	Icon         string `json:"icon"`
	Active       string `json:"active"`
	ChildrenList []Item `json:"children_list"`
}


func constructMenuTrees(menus []models.Menu, parentId int) []Item {
	
	branch := make([]Item, 0)
	
	for  _,menu := range menus {
		if menu.ParentId == parentId{
			childList := constructMenuTrees(menus, menu.Id)
			
			child := Item{
				MenuName:     menu.MenuName,
				ID:           menu.Id,
				Url:          menu.Url,
				Icon:         menu.Icon,
				Active:       "",
				ChildrenList: childList,
			}
			branch = append(branch, child)
		}
	}
	
	return branch
}

func constructMenuTree(menus []map[string]interface{}, parentId int64) []Item {
	
	branch := make([]Item, 0)
	
	for j := 0; j < len(menus); j++ {
		if parentId == menus[j]["parent_id"].(int64) {
			
			childList := constructMenuTree(menus, menus[j]["id"].(int64))
			
			child := Item{
				MenuName:         menus[j]["menu_name"].(string),
				ID:           menus[j]["id"].(int),
				Url:          menus[j]["url"].(string),
				Icon:         menus[j]["icon"].(string),
				Active:       "",
				ChildrenList: childList,
			}
			
			branch = append(branch, child)
		}
	}
	
	return branch
}