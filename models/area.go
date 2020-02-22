package models

import (
	"errors"
)

type Area struct {
	Model
	Id       int      `json:"id"       form:"id"       gorm:"default:''"`
	Adcode   string   `json:"adcode"   form:"adcode"   gorm:"default:''"`
	Citycode int      `json:"citycode" form:"citycode" gorm:"default:''"`
	Center   string   `json:"center"   form:"center"   gorm:"default:''"`
	Name     string   `json:"name"     form:"name"     gorm:"default:''"`
	ParentId int      `json:"parent_id"form:"parent_id"gorm:"default:''"`
	IsEnd    int      `json:"is_end"   form:"is_end"   gorm:"default:'1'"`
	
	Child    []*Area   `gorm:"-"`   // 忽略这个字段
}


func NewArea() (area *Area) {
	return &Area{}
}

func (m *Area) Pagination(offset, limit int, key string) (res []Area, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Area{}).Count(&count)
	return
}

func (m *Area) Create() (newAttr Area, err error) {
	
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

func (m *Area) Update() (newAttr Area, err error) {
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

func (m *Area) Delete() (err error) {
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

func (m *Area) DelBatch(ids []int) (err error) {
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

func (m *Area) FindById(id int) (area Area, err error) {
	err = Db.Where("id=?", id).First(&area).Error
	return
}

func (m *Area) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Area, total int64, err error) {
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
	query.Model(&Area{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

