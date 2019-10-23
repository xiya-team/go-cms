package models

import (
	"errors"
	"time"
)

type Role struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	RoleName  string    `json:"role_name" form:"role_name" gorm:"default:''"`
	RoleKey   string    `json:"role_key"  form:"role_key"  gorm:"default:''"`
	RoleSort  int       `json:"role_sort" form:"role_sort" gorm:"default:''"`
	DataScope int    	`json:"data_scope"form:"data_scope"gorm:"default:'1'"`
	Status    int    	`json:"status"    form:"status"    gorm:"default:''"`
	DelFlag   int    	`json:"del_flag"  form:"del_flag"  gorm:"default:'0'"`
	CreateBy  int    	`json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt time.Time `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  int    	`json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt time.Time `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	
}


func NewRole() (role *Role) {
	return &Role{}
}

func (m *Role) Pagination(offset, limit int, key string) (res []Role, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Role{}).Count(&count)
	return
}

func (m *Role) Create() (newAttr Role, err error) {

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

func (m *Role) Update() (newAttr Role, err error) {
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

func (m *Role) Delete() (err error) {
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

func (m *Role) DelBatch(ids []int) (err error) {
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

func (m *Role) FindById(id int) (role Role, err error) {
	err = Db.Where("id=?", id).First(&role).Error
	return
}

func (m *Role) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Role, total int64, err error) {
	query := Db
	if status,isExist:=dataMap["status"].(int);isExist{
		query = query.Where("status = ?", status)
	}
	if name,ok:=dataMap["name"].(string);ok{
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if startTime,ok:=dataMap["start_time"].(int64);ok{
		query = query.Where("created_at > ?", startTime)
	}
	if endTime,ok:=dataMap["end_time"].(int64);ok{
		query = query.Where("created_at <= ?", endTime)
	}

    if orderBy!=""{
		query = query.Order(orderBy)
	}

	// 获取取指page，指定pagesize的记录
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	if err == nil{
		err = query.Model(&m).Count(&total).Error
	}
	return
}

