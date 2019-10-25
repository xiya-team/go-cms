package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-cms/common"
	"log"
	"sync"
	"time"
)

var Db *gorm.DB

type Model struct {
	StartTime   time.Time   `json:"-" gorm:"-" form:"start_time"`   // 忽略这个字段
	EndTime     time.Time   `json:"-" gorm:"-" form:"end_time"`   // 忽略这个字段
	Page        int64       `json:"-" gorm:"-" form:"page"`   // 忽略这个字段
	PageSize    int64       `json:"-" gorm:"-" form:"page_size"`   // 忽略这个字段
	OrderColumnName  string `json:"-" gorm:"-" form:"order_column_name"`   // 忽略这个字段
	OrderType     string    `json:"-" gorm:"-" form:"order_type"`   // 忽略这个字段
}

func NewModel() (model *Model) {
	return &Model{}
}

const (
	RECODE_OK         = "0"
	RECODE_DBERR      = "4001"
	RECODE_NODATA     = "4002"
	RECODE_DATAEXIST  = "4003"
	RECODE_DATAERR    = "4004"
	RECODE_SESSIONERR = "4101"
	RECODE_LOGINERR   = "4102"
	RECODE_PARAMERR   = "4103"
	RECODE_USERERR    = "4104"
	RECODE_ROLEERR    = "4105"
	RECODE_PWDERR     = "4106"
	RECODE_REQERR     = "4201"
	RECODE_IPERR      = "4202"
	RECODE_THIRDERR   = "4301"
	RECODE_IOERR      = "4302"
	RECODE_SERVERERR  = "4500"
	RECODE_UNKNOWERR  = "4501"
)

var recodeText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
}

func RecodeText(code string) string{
	str,ok := recodeText[code]
	if ok {
		return str
	}
	return RecodeText(RECODE_UNKNOWERR)
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
		OnceMutex sync.Once
	)

	dbType = beego.AppConfig.String("dbType")
	dbName = beego.AppConfig.String("dbName")
	user = beego.AppConfig.String("user")
	password = beego.AppConfig.String("password")
	host = beego.AppConfig.String("host")
	tablePrefix = beego.AppConfig.String("tablePrefix")
	
	//只执行一次
	OnceMutex.Do(func() {
		Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName))
		
		if err != nil {
			log.Println(err)
		}
		
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return tablePrefix + defaultTableName
		}
		
		//开发模式打印SQL
		if beego.BConfig.RunMode == "dev" {
			Db.LogMode(true)
			//同步数据库
			//Db.AutoMigrate(&User{})
		}
		
		Db.SingularTable(true)
		Db.DB().SetMaxIdleConns(10)
		Db.DB().SetMaxOpenConns(100)
		
		Db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
		Db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
		Db.Callback().Delete().Replace("gorm:update_time_stamp",updateTimeStampForDeleteCallback)
		
		Db.Callback().Create().Register("create_admin_log", CreateAdminLogCallback)
		Db.Callback().Update().Register("update_admin_log", UpdateAdminLogCallback)
		//Db.Callback().Update().Remove("gorm:xxx")
		Db.Callback().Delete().Register("delete_admin_log", DeleteAdminLogCallback)
	})
	
}

func CloseDB() {
	defer Db.Close()
}

func GetMysqlMsg() (mysqlMsg map[string]string) {
	mysqlMsg = make(map[string]string)
	var version string
	if err := Db.Raw("select version()").Row().Scan(&version); err != nil {
		log.Println(err)
	}
	mysqlMsg["version"] = version
	return
}

func CreateAdminLogCallback(scope *gorm.Scope) {
	if scope.TableName() != "cms_admin_log" {
		//adminLogModel := NewAdminLog()
		//adminLogModel.CreatedAt = php2go.Time()
		//adminLogModel.UpdatedAt = php2go.Time()
		////adminLogModel.Ip = common.Ctx.Input.IP()
		//adminLogModel.Ip = "127.0.0.1"
		//adminLogModel.UserId = common.UserId
		//adminLogModel.Route = ""
		//adminLogModel.Method = ""
		//adminLogModel.Description = fmt.Sprintf("%s添加了表%s 的%s", common.UserId, scope.TableName(), fmt.Sprintf("%+v", scope.Value))
		//adminLogModel.Create()
		//
		//Db.Create(&AdminLog{
		//	Route: common.Ctx.Request.URL.String(),
		//	UserId:      common.UserId,
		//	Ip:          common.Ctx.Input.IP(),
		//	Method:      common.Ctx.Request.Method,
		//	Description: fmt.Sprintf("%s添加了表%s 的%s", common.UserId, scope.TableName(), fmt.Sprintf("%+v", scope.Value)),
		//})
	}
	return
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	
	if !scope.HasError() {
		nowTime := time.Now()
		if createAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createAtField.IsBlank {
				createAtField.Set(nowTime)
			}
		}
		
		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(nowTime)
			}
		}
	}
}

// 注册更新钩子在持久化之前
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
		if updatedAtField.IsBlank {
			updatedAtField.Set(time.Now())
		}
	}
}

// 注册更新钩子在持久化之前
func updateTimeStampForDeleteCallback(scope *gorm.Scope) {
	if deleteAtField, ok := scope.FieldByName("DeletedAt"); ok {
		if deleteAtField.IsBlank {
			deleteAtField.Set(time.Now())
		}
	}
}

func UpdateAdminLogCallback(scope *gorm.Scope) {
	if common.Ctx != nil {
		//Db.Create(&AdminLog{
			//Route: common.Ctx.Request.URL.String(),
			//UserId:      common.UserId,
			//Ip:          common.Ctx.Input.IP(),
			//Method:      common.Ctx.Request.Method,
			//Description: fmt.Sprintf("%s修改了表%s 的%s", common.UserId, scope.TableName(), fmt.Sprintf("%+v", scope.Value)),
		//})
	}
	return
}

func DeleteAdminLogCallback(scope *gorm.Scope) {
	//Db.Create(&AdminLog{
		//Route: common.Ctx.Request.URL.String(),
		//UserId:      common.UserId,
		//Ip:          common.Ctx.Input.IP(),
		//Method:      common.Ctx.Request.Method,
		//Description: fmt.Sprintf("%s删除了表%s 的一条数据", common.UserId, scope.TableName(), scope.Value),
	//})
	return
}