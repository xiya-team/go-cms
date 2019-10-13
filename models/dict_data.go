package models

import "errors"

type DictData struct {
	Model
	DictCode  int       `json:"dict_code" form:"dict_code" gorm:"default:''"`
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	DictSort  int       `json:"dict_sort" form:"dict_sort" gorm:"default:'0'"`
	DictLabel string    `json:"dict_label"form:"dict_label"gorm:"default:''"`
	DictValue string    `json:"dict_value"form:"dict_value"gorm:"default:''"`
	DictType  string    `json:"dict_type" form:"dict_type" gorm:"default:''"`
	CssClass  string    `json:"css_class" form:"css_class" gorm:"default:''"`
	ListClass string    `json:"list_class"form:"list_class"gorm:"default:''"`
	IsDefault string    `json:"is_default"form:"is_default"gorm:"default:'N'"`
	Status    string    `json:"status"    form:"status"    gorm:"default:'0'"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt int       `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt int       `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	
}


func NewDictData() (dictData *DictData) {
	return &DictData{}
}

func (m *DictData) Pagination(offset, limit int, key string) (res []DictData, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(DictData{}).Count(&count)
	return
}

func (m *DictData) Create() (newAttr DictData, err error) {

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

func (m *DictData) Update() (newAttr DictData, err error) {
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

func (m *DictData) Delete() (err error) {
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

func (m *DictData) DelBatch(ids []int) (err error) {
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

func (m *DictData) FindById(id int) (dictData DictData, err error) {
	err = Db.Where("id=?", id).First(&dictData).Error
	return
}

func (m *DictData) FindByMap(offset, limit int, dataMap map[string]interface{},orderBy string) (res []DictData, total int, err error) {
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
		err = query.Model(&User{}).Count(&total).Error
	}
	return
}

