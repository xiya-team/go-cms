package generate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var TplModel = `package models

import (
	"errors"
	"time"
)

type Config struct {
	Model
	$attr$
}

func NewConfig() (config *Config) {
	return &Config{}
}

func (m *Config) Pagination(offset, limit int, key string) (res []Config, count int) {
	query := Db
	if key != "" {
		query = query.Select("*").Where("name like ?", "%"+key+"%")
	}
	query.Select("*").Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Config{}).Count(&count)
	return
}

func (m *Config) Create() (newAttr Config, err error) {
    tx := Db.Begin()
	err = tx.Create(m).Error
	
	if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}

	newAttr = *m
	return
}

func (m *Config) Update() (newAttr Config, err error) {
    tx := Db.Begin()
	if m.Id > 0 {
        err = tx.Model(&m).Where("id=?", m.Id).Updates(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}
	newAttr = *m
	return
}

func (m *Config) Delete() (err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Model(&m).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}
	return
}

func (m *Config) DelBatch(ids []int) (err error) {
    tx := Db.Begin()
	if len(ids) > 0 {
		err = tx.Model(&m).Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}
	return
}

func (m *Config) FindById(id int) (config Config, err error) {
	err = Db.Select("*").Where("id=?", id).First(&config).Error
	return
}

func (m *Config) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Config, total int64, err error) {
	query := Db
	if status,isExist:=dataMap["status"].(int);isExist{
		query = query.Where("status = ?", status)
	}

	if name,ok:=dataMap["name"].(string);ok{
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if startTime,ok:=dataMap["start_time"].(string);ok{
		query = query.Where("created_at > ?", startTime)
	}
	if endTime,ok:=dataMap["end_time"].(string);ok{
		query = query.Where("created_at <= ?", endTime)
	}
    
    if fields,ok:=dataMap["fields"].(string);ok{
		query = query.Select(fields)
	}else {
		query = query.Select("*")
	}

    if orderBy!=""{
		query = query.Order(orderBy)
	}

	// 获取取指page，指定pagesize的记录
	query.Model(&Config{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

`

type FcModel struct {
}

func CreateModel(modelPath, tableName string, tableAttr []TabelAttr, ) {
	if err := os.MkdirAll(path.Clean(modelPath), 777); err != nil {
		log.Println("模型文件创建失败")
	}
	modelData := CreateModelBase(tableAttr, tableName, modelPath, TplModel)
	if err := ioutil.WriteFile(path.Join(modelPath, fmt.Sprintf("%s.go", tableName)), []byte(modelData), os.ModeType); err != nil {
		log.Println(err)
	}
}

func CreateModelBase(tableAttr []TabelAttr, tableName, modelPath string, tpl string) (tplModel string) {
	var maxLongField int
	for _, v := range tableAttr {
		if len(v.Field) > maxLongField {
			maxLongField = len(v.Field)
		}
	}

	//字段名称
	var attr string
	for _, v := range tableAttr {
		fieldName := Hump(v.Field, "max")
		//类型
		typeName := getType(v.Type)
		f1 := fieldName + strings.Repeat(" ", maxLongField-len(fieldName))
		f2 := typeName + strings.Repeat(" ", maxLongField-len(typeName))
		attr += f1 + f2 + fmt.Sprintf("`json:\"%s\"%sform:\"%s\"%sgorm:\"default:'%s'\"`\n	",
			v.Field,
			strings.Repeat(" ", maxLongField-len(v.Field)),
			v.Field,
			strings.Repeat(" ", maxLongField-len(v.Field)),
			v.Default,
		)
	}

	//替换字符
	tplModel = ReplaceStr(tableName, modelPath, tpl, attr)

	return
}
