package models

import (
	"errors"
	"time"
)

type DictData struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	DictId    int       `json:"dict_id"   form:"dict_id"   gorm:"default:''" validate:"required"`
	DictSort  int       `json:"dict_sort" form:"dict_sort" gorm:"default:'0'"`
	DictLabel string    `json:"dict_label"form:"dict_label"gorm:"default:''"`
	DictValue string    `json:"dict_value"form:"dict_value"gorm:"default:''"`
	DictNumber int      `json:"dict_number"form:"dict_number"gorm:"default:''"`
	DictType  string    `json:"dict_type" form:"dict_type" gorm:"default:''"`
	DictValueType  int  `json:"dict_value_type" form:"dict_value_type" gorm:"default:''"`
	CssClass  string    `json:"css_class" form:"css_class" gorm:"default:''"`
	ListClass string    `json:"list_class"form:"list_class"gorm:"default:''"`
	IsDefault int       `json:"is_default"form:"is_default"gorm:"default:'1'"`
	Status    int       `json:"status"    form:"status"    gorm:"default:'0'"`
	CreateBy  int       `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt time.Time `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  int       `json:"update_by" form:"update_by" gorm:"default:''"`
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

func (m *DictData) FindWhere(dataMap map[string]interface{}) (dictData DictData, err error) {
	query := Db
	if dictId,isExist:=dataMap["dict_id"].(int);isExist{
		query = query.Where("dict_id = ?", dictId)
	}

	if dictValue,isExist:=dataMap["dict_value"].(int);isExist{
		query = query.Where("dict_value = ?", dictValue)
	}

	if dictNumber,isExist:=dataMap["dict_number"].(int);isExist{
		query = query.Where("dict_number = ?", dictNumber)
	}

	err = query.First(&dictData).Error
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

	if dictType,isExist:=dataMap["dict_type"].(string);isExist{
		query = query.Where("dict_type = ?", dictType)
	}

	if dictValueType,isExist:=dataMap["dict_value_type"].(int);isExist{
		query = query.Where("dict_value_type = ?", dictValueType)
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
	}else {
		query = query.Select("*")
	}

    if orderBy!=""{
		query = query.Order(orderBy)
	}

	// 获取取指page，指定pagesize的记录
	query.Model(&DictData{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

