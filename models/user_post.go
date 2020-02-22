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
	
	if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}

	newAttr = *m
	return
}

func (m *UserPost) Update() (newAttr UserPost, err error) {
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

func (m *UserPost) Delete() (err error) {
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

func (m *UserPost) DelBatch(ids []int) (err error) {
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

func (m *UserPost) FindById(id int) (userPost UserPost, err error) {
	err = Db.Where("id=?", id).First(&userPost).Error
	return
}

func (m *UserPost) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []UserPost, total int64, err error) {
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

    if orderBy!=""{
		query = query.Order(orderBy)
	}

	// 获取取指page，指定pagesize的记录
	query.Model(&UserPost{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

