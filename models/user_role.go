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

func (m *UserRole) Delete() (err error) {
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

func (m *UserRole) DelBatch(ids []int) (err error) {
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

func (m *UserRole) FindById(id int) (userRole UserRole, err error) {
	err = Db.Where("id=?", id).First(&userRole).Error
	return
}

func (m *UserRole) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []UserRole, total int64, err error) {
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
	query.Model(&UserRole{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

