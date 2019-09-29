package generate

import (
	"flag"
	"go-cms/models"
	"github.com/astaxie/beego"
	"path/filepath"
	"strings"
)

type TabelAttr struct {
	Field   string `gorm:"column:Field"`
	Type    string `gorm:"column:Type"`
	Null    string `gorm:"column:Null"`
	Key     string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra   string `gorm:"column:Extra"`
}

// sudo go run main.go -m models -c controllers/sys -t menu
// go run main.go -t post -c controllers/blog -m models -v views/index
//从根路径开始，填写文件路径，生成的文件名为数据库名称
func Run() {
	tableName := flag.String("t", "", "表名")
	modelPath := flag.String("m", "", "模型地址")
	controllerPath := flag.String("c", "", "控制器地址")
	viewPath := flag.String("v", "", "视图地址")
	flag.Parse()
	if *tableName != "" {
		Generate(*tableName, *modelPath, *controllerPath, *viewPath)
		return
	}
}

func Generate(tableName, modelPath, controllerPath, viewPath string) {
	tablePrefix := beego.AppConfig.String("tablePrefix")
	var tableAttr []TabelAttr
	models.Db.Raw("desc " + tablePrefix + tableName).Scan(&tableAttr)

	//
	if modelPath != "" {
		CreateModel(modelPath, tableName, tableAttr)
	}
	//
	if controllerPath != "" {
		CreateController(controllerPath, tableName)
	}

}

//字段类型转换
func getType(typeName string) (str string) {
	if strings.Index(typeName, "bigint") >= 0 || strings.Index(typeName, "int") >= 0 || strings.Index(typeName, "tinyint") >= 0 {
		return "int"
	}
	if strings.Index(typeName, "varchar") >= 0 || strings.Index(typeName, "char") >= 0 || strings.Index(typeName, "text") >= 0 {
		return "string"
	}
	if strings.Index(typeName, "datetime") >= 0 || strings.Index(typeName, "time") >= 0 {
		return "time.Time"
	}
	return "string"
}

//字段转换大驼峰，小驼峰
//mysql 自带命名要使用小写"_"分割
func Hump(v, t string) (new string) {
	field := strings.Split(v, "_")
	if t == "min" {
		for k, v := range field {
			if k > 0 {
				field[k] = strings.Title(v)
			}
		}
	}
	if t == "max" {
		for k, v := range field {
			field[k] = strings.Title(v)
		}
	}
	new = strings.Join(field, "")
	return
}

func ReplaceStr(tableName, path, oldStr, option string) (newStr string) {
	newStr = strings.Replace(oldStr, "Category", Hump(tableName, "max"), -1)
	newStr = strings.Replace(newStr, "category", Hump(tableName, "min"), -1)

	_, f := filepath.Split(path)
	newStr = strings.Replace(newStr, "$path$", f, -1)

	newStr = strings.Replace(newStr, "package models", "package "+filepath.Clean(path), -1)
	newStr = strings.Replace(newStr, "Config", Hump(tableName, "max"), -1)
	newStr = strings.Replace(newStr, "config", Hump(tableName, "min"), -1)
	newStr = strings.Replace(newStr, "$attr$", option, -1)
	return
}
