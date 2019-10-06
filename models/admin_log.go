package models

import "errors"

type AdminLog struct {
	Model
	Id         int        `json:"id"         form:"id"         gorm:"default:''"`
	Route      string     `json:"route"      form:"route"      gorm:"default:''"`
	Method     string     `json:"method"     form:"method"     gorm:"default:''"`
	Description string     `json:"description"form:"description"gorm:"default:''"`
	UserId     int        `json:"user_id"    form:"user_id"    gorm:"default:'0'"`
	Ip         int        `json:"ip"         form:"ip"         gorm:"default:'0'"`
	CreatedAt  int        `json:"created_at" form:"created_at" gorm:"default:'0'"`
	UpdatedAt  int        `json:"updated_at" form:"updated_at" gorm:"default:'0'"`
	DeletedAt  int        `json:"deleted_at" form:"deleted_at" gorm:"default:'0'"`
	
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

func (m *AdminLog) Delete() (err error) {
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

func (m *AdminLog) DelBatch(ids []int) (err error) {
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

func (m *AdminLog) FindById(id int) (adminLog AdminLog, err error) {
	err = Db.Where("id=?", id).First(&adminLog).Error
	return
}

