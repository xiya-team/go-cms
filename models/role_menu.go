package models

import "errors"

type RoleMenu struct {
	Model
	RoleId int    `json:"role_id"form:"role_id"gorm:"default:''"`
	MenuId int    `json:"menu_id"form:"menu_id"gorm:"default:''"`
	Id     int    `json:"id"     form:"id"     gorm:"default:''"`
	
}


func NewRoleMenu() (roleMenu *RoleMenu) {
	return &RoleMenu{}
}

func (m *RoleMenu) Pagination(offset, limit int, key string) (res []RoleMenu, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(RoleMenu{}).Count(&count)
	return
}

func (m *RoleMenu) Create() (newAttr RoleMenu, err error) {

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

func (m *RoleMenu) Update() (newAttr RoleMenu, err error) {
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

func (m *RoleMenu) Delete() (err error) {
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

func (m *RoleMenu) DelBatch(ids []int) (err error) {
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

func (m *RoleMenu) FindById(id int) (roleMenu RoleMenu, err error) {
	err = Db.Where("id=?", id).First(&roleMenu).Error
	return
}

func (m *RoleMenu) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []RoleMenu, total int64, err error) {
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

	if fields,ok:=dataMap["fields"].(string);ok{
		query = query.Select(fields)
	}else {
		query = query.Select("*")
	}

    if orderBy!=""{
		query = query.Order(orderBy)
	}

	// 获取取指page，指定pagesize的记录
	query.Model(&RoleMenu{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

