package models

import (
	"errors"
	"time"
)

type DictType struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	DictName  string    `json:"dict_name" form:"dict_name" gorm:"default:''"`
	DictType  string    `json:"dict_type" form:"dict_type" gorm:"default:''"`
	Status    string    `json:"status"    form:"status"    gorm:"default:'0'"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt int       `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt int       `json:"updated_at"form:"updated_at"gorm:"default:''"`
	DeletedAt time.Time   `json:"deleted_at"  form:"deleted_at"  gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	
	StartTime   int64       `form:"start_time"   gorm:"-"`   // 忽略这个字段
	EndTime     int64       `form:"end_time"     gorm:"-"`   // 忽略这个字段
}


func NewDictType() (dictType *DictType) {
	return &DictType{}
}

func (m *DictType) Pagination(offset, limit int, key string) (res []DictType, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(DictType{}).Count(&count)
	return
}

func (m *DictType) Create() (newAttr DictType, err error) {

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

func (m *DictType) Update() (newAttr DictType, err error) {
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

func (m *DictType) Delete() (err error) {
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

func (m *DictType) DelBatch(ids []int) (err error) {
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

func (m *DictType) FindById(id int) (dictType DictType, err error) {
	err = Db.Where("id=?", id).First(&dictType).Error
	return
}

func (m *DictType) FindByMap(offset, limit int, dataMap map[string]interface{},orderBy string) (res []DictType, total int, err error) {
	query := Db
	if status,isExist:=dataMap["status"].(int);isExist{
		query = query.Where("status = ?", status)
	}
	
	if dictName,ok:=dataMap["dict_name"].(string);ok{
		query = query.Where("dict_name LIKE ?", "%"+dictName+"%")
	}
	
	if dictType,ok:=dataMap["dict_type"].(string);ok{
		query = query.Where("dict_type LIKE ?", "%"+dictType+"%")
	}

	if startTime,ok:=dataMap["start_time"].(int64);ok{
		query = query.Where("created_at > ?", startTime)
	}
	if endTime,ok:=dataMap["end_time"].(int64);ok{
		query = query.Where("created_at <= ?", endTime)
	}

    if orderBy!=""{
		query = query.Order(orderBy)
	}

	// 获取取指page，指定pagesize的记录
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	if err == nil{
		err = query.Model(&m).Count(&total).Error
	}
	return
}

