package models

import "errors"

type DictType struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	DictId    int       `json:"dict_id"   form:"dict_id"   gorm:"default:''"`
	DictName  string    `json:"dict_name" form:"dict_name" gorm:"default:''"`
	DictType  string    `json:"dict_type" form:"dict_type" gorm:"default:''"`
	Status    string    `json:"status"    form:"status"    gorm:"default:'0'"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt int       `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt int       `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	
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
	
	if err == nil{
		tx.Commit()
	}else {
		tx.Rollback()
	}

	newAttr = *m
	return
}

func (m *DictType) Update() (newAttr DictType, err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err == nil{
		tx.Commit()
	}else {
		tx.Rollback()
	}
	newAttr = *m
	return
}

func (m *DictType) Delete() (err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err == nil{
		tx.Commit()
	}else {
		tx.Rollback()
	}
	return
}

func (m *DictType) DelBatch(ids []int) (err error) {
    tx := Db.Begin()
	if len(ids) > 0 {
		err = tx.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err == nil{
		tx.Commit()
	}else {
		tx.Rollback()
	}
	return
}

func (m *DictType) FindById(id int) (dictType DictType, err error) {
	err = Db.Where("id=?", id).First(&dictType).Error
	return
}

