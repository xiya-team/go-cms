package sys

import (
	"github.com/astaxie/beego/validation"
	"github.com/syyongx/php2go"
	"go-cms/controllers"
	"go-cms/models"
	"go-cms/pkg/e"
	"go-cms/pkg/util"
	"log"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Prepare() {

}

func (c *UserController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewUser().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "user/index.html"
}

func (c *UserController) Create() {
	if c.Ctx.Input.IsPost() {
		userModel := models.NewUser()
		//1.压入数据
		if err := c.ParseForm(userModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(userModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := userModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.User{}
	c.TplName = c.ADMIN_TPL + "user/create.html"
}

func (c *UserController) Update() {
	id, _ := c.GetInt("id")
	user, _ := models.NewUser().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&user); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(user); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := user.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["vo"] = user
	c.TplName = c.ADMIN_TPL + "user/update.html"
}

func (c *UserController) Delete() {
	userModel := models.NewUser()
	id, _ := c.GetInt("id")
	userModel.Id = id
	if err := userModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *UserController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	userModel := models.NewUser()
	if err := userModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *UserController) Login() {
	user_name := c.GetString("user_name")
	password := c.GetString("password")
	
	u := models.User{UserName: user_name, Password: password}
	
	valid := validation.Validation{}
	valid.Required(u.UserName, "用户名").Message("不能为空!")
	valid.Required(u.Password, "密码").Message("不能为空!")
	
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			c.JsonResult(e.ERROR, err.Key+":"+err.Message)
		}
	}
	
	userModel := models.NewUser()
	user, _ :=userModel.FindByUserName(user_name)
	
	if php2go.Empty(user) {
		c.JsonResult(e.ERROR, "User Not Exist")
	}
	
	has := php2go.Md5(password+user.Salt)
	
	if(user.Password == has){
		token:=util.CreateToken(user)
		jsonData := make(map[string]interface{}, 1)
		jsonData["token"] = token
		c.JsonResult(e.SUCCESS,"登录成功!",jsonData)
	}
	
	c.JsonResult(e.ERROR, has)
}

