package models

import (
	"errors"
	"time"
)

type DictData struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	DictId    int       `json:"dict_id"   form:"dict_id"   gorm:"default:''"`
	DictSort  int       `json:"dict_sort" form:"dict_sort" gorm:"default:'0'"`
	DictLabel string    `json:"dict_label"form:"dict_label"gorm:"default:''"`
	DictValue string    `json:"dict_value"form:"dict_value"gorm:"default:''"`
	DictType  string    `json:"dict_type" form:"dict_type" gorm:"default:''"`
	CssClass  string    `json:"css_class" form:"css_class" gorm:"default:''"`
	ListClass string    `json:"list_class"form:"list_class"gorm:"default:''"`
	IsDefault string    `json:"is_default"form:"is_default"gorm:"default:'N'"`
	Status    string    `json:"status"    form:"status"    gorm:"default:'0'"`
	CreateBy  int    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt time.Time `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  int    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt time.Time `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
}


func NewDictData() (dictData *DictData) {
	return &DictData{}
}

func (m *DictData) Pagination(offset, limit int, key string) (res []DictData, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Select("*").Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(DictData{}).Count(&count)
	return
}

func (m *DictData) Create() (newAttr DictData, err error) {

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

func (m *DictData) Update() (newAttr DictData, err error) {
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

func (m *DictData) Delete() (err error) {
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

func (m *DictData) DelBatch(ids []int) (err error) {
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

func (m *DictData) FindById(id int) (dictData DictData, err error) {
	err = Db.Where("id=?", id).First(&dictData).Error
	return
}

func (m *DictData) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []DictData, total int64, err error) {
	query := Db
	if status,isExist:=dataMap["status"].(int);isExist{
		query = query.Where("status = ?", status)
	}
	
	if dictId,isExist:=dataMap["dict_id"].(int);isExist{
		query = query.Where("dict_id = ?", dictId)
	}
	
	if dictLabel,ok:=dataMap["dict_label"].(string);ok{
		query = query.Where("dict_label LIKE ?", "%"+dictLabel+"%")
	}

	if startTime,ok:=dataMap["start_time"].(string);ok{
		query = query.Where("created_at > ?", startTime)
	}
	if endTime,ok:=dataMap["end_time"].(string);ok{
		query = query.Where("created_at <= ?", endTime)
	}

	if fields,ok:=dataMap["fields"].(string);ok{
		query = query.Select(fields)
	}

    if orderBy!=""{
		query = query.Order(orderBy)
	}

	// 获取取指page，指定pagesize的记录
	err = query.Select("*").Offset(offset).Limit(limit).Find(&res).Count(&total).Error
	return
}

