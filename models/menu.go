package models

import "errors"

type Menu struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	MenuName  string    `json:"menu_name" form:"menu_name" gorm:"default:''"`
	ParentId  int       `json:"parent_id" form:"parent_id" gorm:"default:'0'"`
	OrderNum  int       `json:"order_num" form:"order_num" gorm:"default:'0'"`
	Url       string    `json:"url"       form:"url"       gorm:"default:'#'"`
	MenuType  int       `json:"menu_type" form:"menu_type" gorm:"default:''"`
	Visible   string    `json:"visible"   form:"visible"   gorm:"default:'0'"`
	Perms     string    `json:"perms"     form:"perms"     gorm:"default:''"`
	Icon      string    `json:"icon"      form:"icon"      gorm:"default:'#'"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt int       `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt int       `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	
}


func NewMenu() (menu *Menu) {
	return &Menu{}
}

func (m *Menu) Pagination(offset, limit int, key string) (res []Menu, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Menu{}).Count(&count)
	return
}

func (m *Menu) Create() (newAttr Menu, err error) {
	err = Db.Create(m).Error
	newAttr = *m
	return
}

func (m *Menu) Update() (newAttr Menu, err error) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Menu) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Menu) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Menu) FindById(id int) (menu Menu, err error) {
	err = Db.Where("id=?", id).First(&menu).Error
	return
}

