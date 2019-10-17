package models

import "errors"

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
	Status    string    `json:"status"    form:"status"    gorm:"default:'0'"`
	DelFlag   string    `json:"del_flag"  form:"del_flag"  gorm:"default:'0'"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt int       `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt int       `json:"updated_at"form:"updated_at"gorm:"default:''"`
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

func (m *Dept) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Dept, total int, err error) {
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

