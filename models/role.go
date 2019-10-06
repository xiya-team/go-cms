package models

import "errors"

type Role struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	RoleName  string    `json:"role_name" form:"role_name" gorm:"default:''"`
	RoleKey   string    `json:"role_key"  form:"role_key"  gorm:"default:''"`
	RoleSort  int       `json:"role_sort" form:"role_sort" gorm:"default:''"`
	DataScope string    `json:"data_scope"form:"data_scope"gorm:"default:'1'"`
	Status    string    `json:"status"    form:"status"    gorm:"default:''"`
	DelFlag   string    `json:"del_flag"  form:"del_flag"  gorm:"default:'0'"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt int       `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt int       `json:"updated_at"form:"updated_at"gorm:"default:''"`
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
	
	if err == nil{
		tx.Commit()
	}else {
		tx.Rollback()
	}

	newAttr = *m
	return
}

func (m *Role) Update() (newAttr Role, err error) {
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

func (m *Role) Delete() (err error) {
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

func (m *Role) DelBatch(ids []int) (err error) {
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

func (m *Role) FindById(id int) (role Role, err error) {
	err = Db.Where("id=?", id).First(&role).Error
	return
}

