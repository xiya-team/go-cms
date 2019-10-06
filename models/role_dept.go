package models

import "errors"

type RoleDept struct {
	Model
	RoleId int    `json:"role_id"form:"role_id"gorm:"default:''"`
	DeptId int    `json:"dept_id"form:"dept_id"gorm:"default:''"`
	Id     int    `json:"id"     form:"id"     gorm:"default:''"`
	
}


func NewRoleDept() (roleDept *RoleDept) {
	return &RoleDept{}
}

func (m *RoleDept) Pagination(offset, limit int, key string) (res []RoleDept, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(RoleDept{}).Count(&count)
	return
}

func (m *RoleDept) Create() (newAttr RoleDept, err error) {

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

func (m *RoleDept) Update() (newAttr RoleDept, err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Where("id=?", m.Id).Save(m).Error
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

func (m *RoleDept) Delete() (err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Delete(m).Error
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

func (m *RoleDept) DelBatch(ids []int) (err error) {
    tx := Db.Begin()
	if len(ids) > 0 {
		err = tx.Where("id in (?)", ids).Delete(m).Error
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

func (m *RoleDept) FindById(id int) (roleDept RoleDept, err error) {
	err = Db.Where("id=?", id).First(&roleDept).Error
	return
}

