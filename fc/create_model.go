package fc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var TplModel = `package models

import "errors"

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
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Config{}).Count(&count)
	return
}

func (m *Config) Create() (newAttr Config, err error) {
	err = Db.Create(m).Error
	newAttr = *m
	return
}

func (m *Config) Update() (newAttr Config, err error) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Config) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Config) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Config) FindById(id int) (config Config, err error) {
	err = Db.Where("id=?", id).First(&config).Error
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
