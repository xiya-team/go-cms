package models

import "errors"

type UserRole struct {
	Model
	UserId int    `json:"user_id"form:"user_id"gorm:"default:''"`
	RoleId int    `json:"role_id"form:"role_id"gorm:"default:''"`
	Id     int    `json:"id"     form:"id"     gorm:"default:''"`
	
}


func NewUserRole() (userRole *UserRole) {
	return &UserRole{}
}

func (m *UserRole) Pagination(offset, limit int, key string) (res []UserRole, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(UserRole{}).Count(&count)
	return
}

func (m *UserRole) Create() (newAttr UserRole, err error) {

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

func (m *UserRole) Update() (newAttr UserRole, err error) {
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

func (m *UserRole) Delete() (err error) {
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

func (m *UserRole) DelBatch(ids []int) (err error) {
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

func (m *UserRole) FindById(id int) (userRole UserRole, err error) {
	err = Db.Where("id=?", id).First(&userRole).Error
	return
}

