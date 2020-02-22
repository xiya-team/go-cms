package models

import (
	"errors"
	"time"
)

type LogInfo struct {
	Model
	Id         int        `json:"id"         form:"id"         gorm:"default:''"`
	Level      int        `json:"level"      form:"level"      gorm:"default:'6'"`
	Path       string     `json:"path"       form:"path"       gorm:"default:''"`
	Get        string     `json:"get"        form:"get"        gorm:"default:''"`
	Method     string     `json:"method"     form:"method"     gorm:"default:''"`
	Post       string     `json:"post"       form:"post"       gorm:"default:''"`
	Message    string     `json:"message"    form:"message"    gorm:"default:''"`
	Ip         string     `json:"ip"         form:"ip"         gorm:"default:''"`
	UserAgent  string     `json:"user_agent" form:"user_agent" gorm:"default:''"`
	Referer    string     `json:"referer"    form:"referer"    gorm:"default:''"`
	CreatedBy  int        `json:"created_by" form:"created_by" gorm:"default:'0'"`
	UpdatedBy   int       `json:"updated_by"  form:"updated_by" gorm:"default:'0'"`
	Status     int        `json:"status"     form:"status"     gorm:"default:''"`
	Username   string     `json:"username"   form:"username"   gorm:"default:''"`
	CreateTime time.Time  `json:"create_time"form:"create_time"gorm:"default:''"`
	
}

func NewLogInfo() (logInfo *LogInfo) {
	return &LogInfo{}
}

func (m *LogInfo) Pagination(offset, limit int, key string) (res []LogInfo, count int) {
	query := Db
	if key != "" {
		query = query.Select("*").Where("name like ?", "%"+key+"%")
	}
	query.Select("*").Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(LogInfo{}).Count(&count)
	return
}

func (m *LogInfo) Create() (newAttr LogInfo, err error) {
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

func (m *LogInfo) Update() (newAttr LogInfo, err error) {
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

func (m *LogInfo) Delete() (err error) {
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

func (m *LogInfo) DelBatch(ids []int) (err error) {
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

func (m *LogInfo) FindById(id int) (logInfo LogInfo, err error) {
	err = Db.Select("*").Where("id=?", id).First(&logInfo).Error
	return
}

func (m *LogInfo) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []LogInfo, total int64, err error) {
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
	query.Model(&LogInfo{}).Count(&total)
	err = query.Select("*").Offset(offset).Limit(limit).Find(&res).Error
	return
}

