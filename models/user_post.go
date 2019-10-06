package models

import "errors"

type UserPost struct {
	Model
	UserId int    `json:"user_id"form:"user_id"gorm:"default:''"`
	PostId int    `json:"post_id"form:"post_id"gorm:"default:''"`
	Id     int    `json:"id"     form:"id"     gorm:"default:''"`
	
}


func NewUserPost() (userPost *UserPost) {
	return &UserPost{}
}

func (m *UserPost) Pagination(offset, limit int, key string) (res []UserPost, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(UserPost{}).Count(&count)
	return
}

func (m *UserPost) Create() (newAttr UserPost, err error) {

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

func (m *UserPost) Update() (newAttr UserPost, err error) {
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

func (m *UserPost) Delete() (err error) {
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

func (m *UserPost) DelBatch(ids []int) (err error) {
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

func (m *UserPost) FindById(id int) (userPost UserPost, err error) {
	err = Db.Where("id=?", id).First(&userPost).Error
	return
}

