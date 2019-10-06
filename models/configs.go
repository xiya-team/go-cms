package models

import "errors"

type Configs struct {
	Model
	Id          int         `json:"id"          form:"id"          gorm:"default:''"`
	ConfigName  string      `json:"config_name" form:"config_name" gorm:"default:''"`
	ConfigKey   string      `json:"config_key"  form:"config_key"  gorm:"default:''"`
	ConfigValue string      `json:"config_value"form:"config_value"gorm:"default:''"`
	ConfigType  string      `json:"config_type" form:"config_type" gorm:"default:'N'"`
	CreatedBy   int         `json:"created_by"  form:"created_by"  gorm:"default:''"`
	UpdatedBy   int         `json:"updated_by"  form:"updated_by"  gorm:"default:''"`
	CreatedAt   int         `json:"created_at"  form:"created_at"  gorm:"default:''"`
	UpdatedAt   int         `json:"updated_at"  form:"updated_at"  gorm:"default:''"`
	DeletedAt   int         `json:"deleted_at"  form:"deleted_at"  gorm:"default:''"`
	Remark      string      `json:"remark"      form:"remark"      gorm:"default:''"`
	
}


func NewConfigs() (configs *Configs) {
	return &Configs{}
}

func (m *Configs) Pagination(offset, limit int, key string) (res []Configs, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Configs{}).Count(&count)
	return
}

func (m *Configs) Create() (newAttr Configs, err error) {

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

func (m *Configs) Update() (newAttr Configs, err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Where("id=?", m.Id).Save(m).Error
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

func (m *Configs) Delete() (err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Delete(m).Error
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

func (m *Configs) DelBatch(ids []int) (err error) {
    tx := Db.Begin()
	if len(ids) > 0 {
		err = tx.Where("id in (?)", ids).Delete(m).Error
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

func (m *Configs) FindById(id int) (configs Configs, err error) {
	err = Db.Where("id=?", id).First(&configs).Error
	return
}

