package models

import (
	"fmt"
	"go-cms/common"
	"go-cms/pkg/str"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
	"sync"
)

var Db *gorm.DB

type Model struct {
	Id        int       `json:"id" form:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" sql:"index"`
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
		
		//Db.LogMode(true)
		Db.SingularTable(true)
		Db.DB().SetMaxIdleConns(10)
		Db.DB().SetMaxOpenConns(100)
		
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
		fmt.Println(scope)
		Db.Create(&AdminLog{Route: common.Fc.Request.URL.String(),
			UserId:      common.UserId,
			Ip:          int(str.Ip2long(common.Fc.Input.IP())),
			Method:      common.Fc.Request.Method,
			Description: fmt.Sprintf("%s添加了表%s 的%s", common.UserId, scope.TableName(), fmt.Sprintf("%+v", scope.Value)),
		})
	}
	return
}

func UpdateAdminLogCallback(scope *gorm.Scope) {
	if common.Fc != nil {
		Db.Create(&AdminLog{Route: common.Fc.Request.URL.String(),
			UserId:      common.UserId,
			Ip:          int(str.Ip2long(common.Fc.Input.IP())),
			Method:      common.Fc.Request.Method,
			Description: fmt.Sprintf("%s修改了表%s 的%s", common.UserId, scope.TableName(), fmt.Sprintf("%+v", scope.Value)),
		})
	}
	return
}

func DeleteAdminLogCallback(scope *gorm.Scope) {
	Db.Create(&AdminLog{Route: common.Fc.Request.URL.String(),
		UserId:      common.UserId,
		Ip:          int(str.Ip2long(common.Fc.Input.IP())),
		Method:      common.Fc.Request.Method,
		Description: fmt.Sprintf("%s删除了表%s 的一条数据", common.UserId, scope.TableName(), scope.Value),
	})
	return
}
