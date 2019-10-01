package models

import "errors"

type User struct {
	Model
	Id         int        `json:"id"         form:"id"         gorm:"default:''"`
	LoginName  string     `json:"login_name" form:"login_name" gorm:"default:''"`
	UserName   string     `json:"user_name"  form:"user_name"  gorm:"default:''"`
	UserType   string     `json:"user_type"  form:"user_type"  gorm:"default:'00'"`
	Email      string     `json:"email"      form:"email"      gorm:"default:''"`
	Phone      string     `json:"phone"      form:"phone"      gorm:"default:''"`
	Phonenumber string     `json:"phonenumber"form:"phonenumber"gorm:"default:''"`
	Sex        string     `json:"sex"        form:"sex"        gorm:"default:'0'"`
	Avatar     string     `json:"avatar"     form:"avatar"     gorm:"default:''"`
	Password   string     `json:"password"   form:"password"   gorm:"default:''"`
	Salt       string     `json:"salt"       form:"salt"       gorm:"default:''"`
	Status     string     `json:"status"     form:"status"     gorm:"default:'0'"`
	DelFlag    string     `json:"del_flag"   form:"del_flag"   gorm:"default:'0'"`
	LoginIp    string     `json:"login_ip"   form:"login_ip"   gorm:"default:''"`
	LoginDate  int        `json:"login_date" form:"login_date" gorm:"default:''"`
	CreateBy   string     `json:"create_by"  form:"create_by"  gorm:"default:''"`
	CreatedAt  int        `json:"created_at" form:"created_at" gorm:"default:''"`
	UpdateBy   string     `json:"update_by"  form:"update_by"  gorm:"default:''"`
	UpdatedAt  int        `json:"updated_at" form:"updated_at" gorm:"default:''"`
	DeletedAt  int        `json:"deleted_at" form:"deleted_at" gorm:"default:''"`
	Remark     string     `json:"remark"     form:"remark"     gorm:"default:''"`
	
}


func NewUser() (user *User) {
	return &User{}
}

func (m *User) Pagination(offset, limit int, key string) (res []User, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(User{}).Count(&count)
	return
}

func (m *User) Create() (newAttr User, err error) {
	err = Db.Create(m).Error
	newAttr = *m
	return
}

func (m *User) Update() (newAttr User, err error) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *User) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *User) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *User) FindById(id int) (user User, err error) {
	err = Db.Where("id=?", id).First(&user).Error
	return
}

/*****************************************************************新增加的方法*****************************************************************/

func (m *User) FindByUserName(user_name string) (user User, err error) {
	err = Db.Select("id,user_name,password,salt").Where("user_name=?", user_name).First(&user).Error
	return
}