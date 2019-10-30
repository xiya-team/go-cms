package sys

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/syyongx/php2go"
	"go-cms/controllers"
	"go-cms/models"
	"go-cms/pkg/e"
	"go-cms/pkg/util"
	"go-cms/services"
	"go-cms/validations/backend"
	"log"
	"strings"
	"time"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Prepare() {

}

/**
获取列表数据
*/
func (c *UserController) Index() {
	if c.Ctx.Input.IsPost() {
		model := models.NewUser()
		
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		err := json.Unmarshal(data, &model)
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		dataMap := make(map[string]interface{}, 0)
		
		if !php2go.Empty(model.Nickname) {
			dataMap["nickname"] = model.Nickname
		}
		if !php2go.Empty(model.UserName) {
			dataMap["user_name"] = model.UserName
		}
		
		if !php2go.Empty(model.Phone) {
			dataMap["phone"] = model.Phone
		}

		if !php2go.Empty(model.DeptId) {
			dataMap["dept_id"] = model.DeptId
		}
		
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
func (c *UserController) Create() {
	if c.Ctx.Input.IsPost() {
		model := models.NewUser()
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
		
		if !php2go.Empty(model.Password) {
			salt := util.Krand(5, 2)
			model.Salt = salt
			model.Password = php2go.Md5(model.Password + salt)
		}
		
		//3.插入数据
		if _, err := model.Create(); err != nil {
			c.JsonResult(e.ERROR, "创建失败")
		}
		c.JsonResult(e.SUCCESS, "添加成功")
	}
}

func (c *UserController) Password()  {
	model := models.NewUser()
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

		if php2go.Md5(model.Password + post.Salt) != post.Password{
			c.JsonResult(e.ERROR, "验证失败，原密码错误。")
		}

		if !php2go.Empty(model.NewPassword) {
			salt := util.Krand(5, 2)
			model.Salt = salt
			model.Password = php2go.Md5(model.NewPassword + salt)
		}

		if _, err := model.Update(); err != nil {
			c.JsonResult(e.ERROR, "修改失败")
		}
		c.JsonResult(e.SUCCESS, "修改成功")
	}
}

/**
更新数据
*/
func (c *UserController) Update() {
	model := models.NewUser()
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
		
		if !php2go.Empty(model.Password) {
			salt := util.Krand(5, 2)
			model.Salt = salt
			model.Password = php2go.Md5(model.Password + salt)
		}
		
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
func (c *UserController) Delete() {
	if c.Ctx.Input.IsDelete() {
		model := models.NewUser()
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

func (c *UserController) BatchDelete() {
	model := models.NewUser()
	
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}
	
	if err := model.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

//用户登录
func (c *UserController) Login() {
	if c.Ctx.Input.IsPost() {
		model := models.NewUser()
		data := c.Ctx.Input.RequestBody
		//json数据封装到user对象中
		
		err := json.Unmarshal(data, model)
		
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		//数据校验
		loginData := backend.UserLoginValidation{}
		loginData.UserName = model.UserName
		loginData.Password = model.Password
		c.ValidatorAuto(&loginData)
		
		//通过service查询
		user := services.FindByUserName(model.UserName)
		
		jsonRes, err := json.Marshal(map[string]interface{}{"Id": user.Id, "UserName": user.UserName})
		if err != nil {
			panic(err)
		}
		
		redisClient := util.NewRedisClient()
		if err != nil{
			c.JsonResult(e.ERROR, "用户名或密码错误!")
		}
		
		err = redisClient.Set("token_"+user.UserName,string(jsonRes),time.Hour*10).Err()
		if err != nil {
			c.JsonResult(e.ERROR, "用户名或密码错误!")
		}
		
		if php2go.Empty(user) {
			c.JsonResult(e.ERROR, "用户名不存在!")
		}
		
		has := php2go.Md5(model.Password + user.Salt)
		
		if (user.Password == has) {
			token := util.CreateToken(user)
			jsonData := make(map[string]interface{}, 4)
			jsonData["token"] = token
			jsonData["userId"] = user.Id
			jsonData["userName"] = user.UserName
			jsonData["nickname"] = user.Nickname
			c.JsonResult(e.SUCCESS, "登录成功!", jsonData)
		}else {
			c.JsonResult(e.ERROR, "用户名或密码错误!")
		}
	}
}

func (c *UserController) Logout()  {
	redisClient := util.NewRedisClient()
	tokenString := c.Ctx.Input.Header(beego.AppConfig.String("jwt::token_name"))
	username := util.GetUserNameByToken(tokenString)
	
	redisClient.Del("token_"+username)
	
	c.JsonResult(e.SUCCESS, "success")
}

func (c *UserController) CheckToken() {
	token := c.Ctx.Input.Header("Authorization")
	
	b, message , code := util.CheckToken(token)
	
	if !b {
		c.JsonResult(code, message)
	}
	
	jsonData := make(map[string]interface{}, 1)
	jsonData["user_id"] = code
	
	c.JsonResult(e.SUCCESS, "success",jsonData)
}

func (c *UserController) UserInfo() {
	token := c.Ctx.Input.Header(beego.AppConfig.String("jwt::token_name"))
	uid := util.GetUserIdByToken(token)
	userInfo, err := models.NewUser().FindById(uid)
	if err != nil {
		c.JsonResult(e.ERROR, e.ResponseMap[e.ERROR])
	}
	c.JsonResult(e.SUCCESS, e.ResponseMap[e.SUCCESS], userInfo)
}