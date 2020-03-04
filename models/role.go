package models

import (
	"errors"
	"github.com/xiya-team/helpers"
	"strconv"
	"strings"
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
	RoleMenu  string    `json:"role_menu" form:"role_menu" gorm:"-"`   // 忽略这个字段
	RoleDept  string    `json:"role_dept" form:"role_dept" gorm:"-"`   // 忽略这个字段
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

	rm := NewRoleMenu()
	if !helpers.Empty(m.RoleMenu) {
		err = tx.Model(&rm).Where("role_id=?", m.Id).Delete(rm).Error
		if err == nil{
			rms := strings.Split(m.RoleMenu, ",")
			for _, value := range rms {
				menu_id, _ := strconv.Atoi(value)
				role_menu := RoleMenu{
					RoleId:m.Id,
					MenuId:menu_id,
				}
				_ = tx.Model(&rm).Create(role_menu).Error
			}
		}
	}

	rd := NewRoleDept()
	if !helpers.Empty(m.RoleDept) {
		err = tx.Model(&rd).Where("role_id=?", m.Id).Delete(rd).Error
		if err == nil{
			rds := strings.Split(m.RoleDept, ",")
			for _, value := range rds {
				dept_id, _ := strconv.Atoi(value)
				role_dept := RoleDept{
					RoleId:m.Id,
					DeptId:dept_id,
				}
				_ = tx.Model(&rd).Create(role_dept).Error
			}
		}
	}
	
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
		rm := NewRoleMenu()
		if !helpers.Empty(m.RoleMenu) {
			err = tx.Model(&rm).Where("role_id=?", m.Id).Delete(rm).Error
			if err == nil{
				rms := strings.Split(m.RoleMenu, ",")
				for _, value := range rms {
					menu_id, _ := strconv.Atoi(value)
					role_menu := RoleMenu{
						RoleId:m.Id,
						MenuId:menu_id,
					}
					_ = tx.Model(&rm).Create(role_menu).Error
				}
			}
		}

		rd := NewRoleDept()
		if !helpers.Empty(m.RoleDept) {
			err = tx.Model(&rd).Where("role_id=?", m.Id).Delete(rd).Error
			if err == nil{
				rds := strings.Split(m.RoleDept, ",")
				for _, value := range rds {
					dept_id, _ := strconv.Atoi(value)
					role_dept := RoleDept{
						RoleId:m.Id,
						DeptId:dept_id,
					}
					_ = tx.Model(&rd).Create(role_dept).Error
				}
			}
		}

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

	rm := NewRoleMenu()
	role_menus := []RoleMenu{}
	err = Db.Model(&rm).Where("role_id=?", m.Id).Find(&role_menus).Error
	if err == nil {
		var role_menu []string
		for _, value := range role_menus {
			role_menu=append(role_menu,strconv.Itoa(value.MenuId))
		}
		role.RoleMenu = strings.Join(role_menu,",")
	}

	rd := NewRoleDept()
	role_depts := []RoleDept{}
	err = Db.Model(&rd).Where("role_id=?", m.Id).Find(&role_depts).Error
	if err == nil {
		var role_dept []string
		for _, value := range role_depts {
			role_dept = append(role_dept,strconv.Itoa(value.DeptId))
		}
		role.RoleDept = strings.Join(role_dept,",")
	}

	return
}

func (m *Role) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Role, total int64, err error) {
	query := Db
	if status,isExist:=dataMap["status"].(int);isExist{
		query = query.Where("status = ?", status)
	}

	if role_name,ok:=dataMap["role_name"].(string);ok{
		query = query.Where("role_name LIKE ?", "%"+role_name+"%")
	}

	//role_key
	if role_key,ok:=dataMap["role_key"].(string);ok{
		query = query.Where("role_key LIKE ?", "%"+role_key+"%")
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
	query.Model(&Role{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

