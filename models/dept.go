package models

import (
	"errors"
	"github.com/wxnacy/wgo/arrays"
	"time"
)

type Dept struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	ParentId  int       `json:"parent_id" form:"parent_id" gorm:"default:'0'"`
	Ancestors string    `json:"ancestors" form:"ancestors" gorm:"default:''"`
	DeptName  string    `json:"dept_name" form:"dept_name" gorm:"default:''"`
	OrderNum  int       `json:"order_num" form:"order_num" gorm:"default:'0'"`
	Leader    string    `json:"leader"    form:"leader"    gorm:"default:''"`
	Phone     string    `json:"phone"     form:"phone"     gorm:"default:''"`
	Email     string    `json:"email"     form:"email"     gorm:"default:''"`
	Status    int       `json:"status"    form:"status"    gorm:"default:'0'"`
	DelFlag   int       `json:"del_flag"  form:"del_flag"  gorm:"default:'0'"`
	CreateBy  int       `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt time.Time `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  int       `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt time.Time `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	
}


func NewDept() (dept *Dept) {
	return &Dept{}
}

func (m *Dept) Pagination(offset, limit int, key string) (res []Dept, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Dept{}).Count(&count)
	return
}

func (m *Dept) Create() (newAttr Dept, err error) {

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

func (m *Dept) Update() (newAttr Dept, err error) {
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

func (m *Dept) Delete() (err error) {
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

func (m *Dept) FindByParentId(parent_id int) (dept Dept, err error) {
	err = Db.Select("*").Where("parent_id=?", parent_id).First(&dept).Error
	return
}

func (m *Dept) DelBatch(ids []int) (err error) {
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

func (m *Dept) FindById(id int) (dept Dept, err error) {
	err = Db.Where("id=?", id).First(&dept).Error
	return
}

func (m *Dept) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Dept, total int64, err error) {
	query := Db
	if status,isExist:=dataMap["status"].(int);isExist{
		query = query.Where("status = ?", status)
	}
	if dept_name,ok:=dataMap["dept_name"].(string);ok{
		query = query.Where("dept_name LIKE ?", "%"+dept_name+"%")
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
	query.Model(&Dept{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}


func (m *Dept) FindAll() (res []Dept, err error) {
	query := Db
	err = query.Find(&res).Error
	return
}

func (m *Dept) FindAllByParentId(parentId int) (res []Dept, err error)   {
	query := Db

	query = query.Where("parent_id = ?", parentId)
	err = query.Find(&res).Error

	return
}

func (m *Dept)FindAllChildren(pid int)  []int {
	var ids []int;
	deptData,_ := m.FindAll()

	for _, dept := range deptData {
		if pid == dept.Id{
			ids = append(ids, dept.Id)
		} else {
			is_exist := arrays.Contains(ids, dept.ParentId)
			if is_exist != -1 {
				ids = append(ids, dept.Id)
			}
		}
	}
	return ids
}
