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

type PaginationRequest struct {
	PageCount int   `form:"page_count" json:"page_count"`
	Page      int   `form:"page" json:"page"`
	StartTime int64 `form:"start_time" json:"start_time"`
	EndTime   int64 `form:"end_time" json:"end_time"`
}
type UserController struct {
	controllers.BaseController
}

func (c *UserController) Prepare() {

}
func (c *UserController) UserList() {
	req := struct {
		PaginationRequest
		models.User
	}{}
	type RespUser struct {
		List  []models.User `json:"list"`
		Count int           `json:"count"`
	}
	err := c.ParseForm(&req)
	if err != nil {
		c.JsonResult(e.ERROR_CODE__JSON__PARSE_FAILED, e.ResponseMap[e.ERROR_CODE__JSON__PARSE_FAILED])
	}
	dataMap := make(map[string]interface{}, 0)
	if req.Status == 1 || req.Status == 2 {
		dataMap["status"] = req.Status
	}
	if req.LoginName != "" {
		dataMap["login_name"] = req.LoginName
	}
	if req.UserName != "" {
		dataMap["user_name"] = req.UserName
	}
	if req.StartTime != 0 {
		dataMap["start_time"] = req.StartTime
	}
	if req.EndTime != 0 {
		dataMap["end_time"] = req.EndTime
	}
	if req.Phone != "" {
		dataMap["phone"] = req.Phone
	}
	page, pageCount := util.InitPageCount(req.Page, req.PageCount)
	var orderBy string = "created_at DESC"
	user, total, err := models.NewUser().FindByMaps(page, pageCount, dataMap, orderBy)
	if err != nil {
		c.JsonResult(e.ERROR, e.ResponseMap[e.ERROR])
	}
	resp := RespUser{
		List:  make([]models.User, 0),
		Count: total,
	}
	if len(user) != 0 {
		resp.List = user
	}
	c.JsonResult(e.SUCCESS, e.ResponseMap[e.SUCCESS], resp)
}
func (c *UserController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewUser().Pagination((page-1)*limit, limit, key)
		c.JsonResult(e.SUCCESS, "ok", result, count)
	}
}
func (c *UserController) UserInfo() {
	token := c.Ctx.Input.Header(beego.AppConfig.String("jwt::token_name"))
	kv := strings.Split(token, " ")
	uid := util.GetUserIdByToken(kv[1])
	userInfo, err := models.NewUser().FindById(uid)
	if err != nil {
		c.JsonResult(e.ERROR, e.ResponseMap[e.ERROR])
	}
	c.JsonResult(e.SUCCESS, e.ResponseMap[e.SUCCESS], userInfo)
}
func (c *UserController) Create() {
	if c.Ctx.Input.IsPost() {
		userModel := models.NewUser()
		//1.压入数据
		if err := c.ParseForm(userModel); err != nil {
			c.JsonResult(e.ERROR, "赋值失败")
		}

		salt := util.Krand(5, 2)
		userModel.Salt = salt
		userModel.LoginName = userModel.UserName
		userModel.Email = userModel.Email
		userModel.Password = php2go.Md5(userModel.Password + salt)
		userModel.CreatedAt = php2go.Time()
		userModel.UpdatedAt = php2go.Time()

		//2.验证
		valid := validation.Validation{}
		b, err := valid.Valid(userModel)
		if err != nil {
			log.Printf("%v\n%v", err, valid.Errors)
			c.JsonResult(e.ERROR, "验证失败")
		}
		if !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
		}

		//3.插入数据
		if _, err := userModel.Create(); err != nil {
			c.JsonResult(e.ERROR, "创建失败")
		}
		c.JsonResult(e.SUCCESS, "添加成功")
	}
}

func (c *UserController) Update() {
	id, _ := c.GetInt("id")
	user, _ := models.NewUser().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&user); err != nil {
			c.JsonResult(e.ERROR, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(user); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(e.ERROR, "验证失败")
		}
		//3
		if _, err := user.Update(); err != nil {
			c.JsonResult(e.ERROR, "修改失败")
		}
		c.JsonResult(e.SUCCESS, "修改成功")
	}

	c.Data["vo"] = user
	c.TplName = c.ADMIN_TPL + "user/update.html"
}

func (c *UserController) Delete() {
	userModel := models.NewUser()
	id, _ := c.GetInt("id")
	userModel.Id = id
	if err := userModel.Delete(); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

func (c *UserController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(e.ERROR, "赋值失败")
	}

	userModel := models.NewUser()
	if err := userModel.DelBatch(ids); err != nil {
		c.JsonResult(e.ERROR, "删除失败")
	}
	c.JsonResult(e.SUCCESS, "删除成功")
}

//用户登录
func (c *UserController) Login() {
	if c.Ctx.Input.IsPost() {
		user_name := c.GetString("user_name")
		password := c.GetString("password")

		//数据校验
		loginData := backend.UserLoginValidation{}
		loginData.UserName = user_name
		loginData.Password = password
		c.ValidatorAuto(&loginData)

		//通过service查询
		user := services.FindByUserName(user_name)
		
		jsonRes, err := json.Marshal(map[string]interface{}{"Id": user.Id, "UserName": user.UserName})
		if err != nil {
			panic(err)
		}
		
		redisClient := util.NewRedisClient()
		if err != nil{
			c.JsonResult(e.ERROR, err.Error())
		}
		
		err = redisClient.Set("token_"+user.UserName,string(jsonRes),time.Hour*10).Err()
		if err != nil {
			c.JsonResult(e.ERROR, err.Error())
		}
		
		if php2go.Empty(user) {
			c.JsonResult(e.ERROR, "User Not Exist")
		}

		has := php2go.Md5(password + user.Salt)

		if (user.Password == has) {
			token := util.CreateToken(user)
			jsonData := make(map[string]interface{}, 1)
			jsonData["token"] = token
			c.JsonResult(e.SUCCESS, "登录成功!", jsonData)
		}

		c.JsonResult(e.ERROR, has)
	}
}

func (c *UserController) CheckToken() {

	token := c.Ctx.Input.Header("Authorization")

	b, _ := util.CheckToken(token)

	if !b {
		c.JsonResult(e.ERROR, "验证失败!")
	}

	c.JsonResult(e.SUCCESS, "success")
}
