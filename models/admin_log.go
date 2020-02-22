package models

import (
	"errors"
	"time"
)

type AdminLog struct {
	Model
	Id         int        `json:"id"         form:"id"         gorm:"default:''"`
	Route      string     `json:"route"      form:"route"      gorm:"default:''"`
	Method     string     `json:"method"     form:"method"     gorm:"default:''"`
	Description string    `json:"description"form:"description"gorm:"default:''"`
	UserId     int        `json:"user_id"    form:"user_id"    gorm:"default:'0'"`
	Ip         string     `json:"ip"         form:"ip"         gorm:"default:'0'"`
	CreatedAt  time.Time  `json:"created_at" form:"created_at" gorm:"default:'0'"`
	UpdatedAt  time.Time  `json:"updated_at" form:"updated_at" gorm:"default:'0'"`
	DeletedAt  time.Time  `json:"deleted_at" form:"deleted_at" gorm:"default:'0'"`
	
}


func NewAdminLog() (adminLog *AdminLog) {
	return &AdminLog{}
}

func (m *AdminLog) Pagination(offset, limit int, key string) (res []AdminLog, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(AdminLog{}).Count(&count)
	return
}

func (m *AdminLog) Create() (newAttr AdminLog, err error) {

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

func (m *AdminLog) Update() (newAttr AdminLog, err error) {
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

func (m *AdminLog) Delete() (err error) {
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

func (m *AdminLog) DelBatch(ids []int) (err error) {
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

func (m *AdminLog) FindById(id int) (adminLog AdminLog, err error) {
	err = Db.Where("id=?", id).First(&adminLog).Error
	return
}

func (m *AdminLog) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []AdminLog, total int64, err error) {
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
	query.Model(&AdminLog{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

