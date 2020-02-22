package models

import (
	"errors"
	"time"
)

type Configs struct {
	Model
	Id          int         `json:"id"          form:"id"          gorm:"default:''"`
	ConfigName  string      `json:"config_name" form:"config_name" gorm:"default:''"`
	ConfigKey   string      `json:"config_key"  form:"config_key"  gorm:"default:''"`
	ConfigValue string      `json:"config_value"form:"config_value" gorm:"default:''"`
	ConfigType  int         `json:"config_type" form:"config_type" gorm:"default:''"`
	CreatedBy   int         `json:"created_by"  form:"created_by"  gorm:"default:''"`
	UpdatedBy   int         `json:"updated_by"  form:"updated_by"  gorm:"default:''"`
	Type   int         `json:"type"  form:"type"  gorm:"default:''"`
	CreatedAt   time.Time   `json:"created_at"  form:"created_at"  gorm:"default:''"`
	UpdatedAt   time.Time   `json:"updated_at"  form:"updated_at"  gorm:"default:''"`
	DeletedAt   time.Time   `json:"deleted_at"  form:"deleted_at"  gorm:"default:''"`
	Remark      string      `json:"remark"      form:"remark"      gorm:"default:''"`
}


func NewConfigs() (configs *Configs) {
	return &Configs{}
}

func (m *Configs) Pagination(offset, limit int, key string) (res []Configs, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Configs{}).Count(&count)
	return
}

func (m *Configs) Create() (newAttr Configs, err error) {
	
	m.CreatedAt = time.Now()
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

func (m *Configs) Update() (newAttr Configs, err error) {
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

func (m *Configs) Delete() (err error) {
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

func (m *Configs) DelBatch(ids []int) (err error) {
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

func (m *Configs) FindById(id int) (configs Configs, err error) {
	err = Db.Where("id=?", id).First(&configs).Error
	return
}

func (m *Configs) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Configs, total int64, err error) {
	query := Db
	if config_type,isExist:=dataMap["config_type"].(int);isExist{
		query = query.Where("config_type = ?", config_type)
	}

	if config_name,ok:=dataMap["config_name"].(string);ok{
		query = query.Where("config_name LIKE ?", "%"+config_name+"%")
	}

	if config_key,ok:=dataMap["config_key"].(string);ok{
		query = query.Where("config_key LIKE ?", "%"+config_key+"%")
	}

	if config_name,ok:=dataMap["config_name"].(string);ok{
		query = query.Where("config_name LIKE ?", "%"+config_name+"%")
	}

	if config_value,ok:=dataMap["config_value"].(string);ok{
		query = query.Where("config_value LIKE ?", "%"+config_value+"%")
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
	query.Model(&Configs{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

func (m *Configs) FindAll(dataMap map[string]interface{}) (res []Configs) {
	query := Db

	if is_type,isExist:=dataMap["type"].(int);isExist{
		query = query.Where(Configs{Type:is_type})
	}

	query.Find(&res)
	return
}