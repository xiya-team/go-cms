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

func (m *RoleDept) Delete() (err error) {
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

func (m *RoleDept) DelBatch(ids []int) (err error) {
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

func (m *RoleDept) FindById(id int) (roleDept RoleDept, err error) {
	err = Db.Where("id=?", id).First(&roleDept).Error
	return
}

func (m *RoleDept) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []RoleDept, total int64, err error) {
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
	query.Model(&RoleDept{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

