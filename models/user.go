package models

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

type User struct {
	Model
	Id          int    `json:"id"         form:"id"         gorm:"default:''"`
	Nickname   string  `json:"nickname"   form:"nickname" gorm:"default:''" valid:"Required;MaxSize(20);MinSize(2)"`
	UserName    string `json:"user_name"  form:"user_name"  gorm:"default:''" valid:"Required;MaxSize(20);MinSize(6)"`
	UserType    int    `json:"user_type"  form:"user_type"  gorm:"default:'00'"`
	Email       string `json:"email"      form:"email"      gorm:"default:''" valid:"Email"`
	Phone       string `json:"phone"      form:"phone"      gorm:"default:''"`
	Phonenumber string `json:"phonenumber"form:"phonenumber"gorm:"default:''"`
	Sex         int    `json:"sex"        form:"sex"        gorm:"default:'1'"`
	Avatar      string `json:"avatar"     form:"avatar"     gorm:"default:''"`
	Password    string `json:"password"   form:"password"   gorm:"default:''" valid:"Required;MaxSize(33);MinSize(6)"`
	Salt        string `json:"salt"       form:"salt"       gorm:"default:''"`
	Status      int    `json:"status"     form:"status"     gorm:"default:'1'"`
	DelFlag     int    `json:"del_flag"   form:"del_flag"   gorm:"default:'1'"`
	LoginIp     string `json:"login_ip"   form:"login_ip"   gorm:"default:''"`
	LoginDate   int64  `json:"login_date" form:"login_date" gorm:"default:''"`
	CreateBy    string `json:"create_by"  form:"create_by"  gorm:"default:''"`
	CreatedAt   int64  `json:"created_at" form:"created_at" gorm:"default:''"`
	UpdateBy    string `json:"update_by"  form:"update_by"  gorm:"default:''"`
	UpdatedAt   int64  `json:"updated_at" form:"updated_at" gorm:"default:''"`
	DeletedAt   time.Time  `json:"deleted_at" form:"deleted_at" gorm:"default:''"`
	Remark      string `json:"remark"     form:"remark"     gorm:"default:''"`
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
	tx := Db.Begin()

	err = Db.Create(m).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	newAttr = *m
	return
}

func (m *User) Update() (newAttr User, err error) {
	tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Model(&m).Where("id=?", m.Id).Updates(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	newAttr = *m
	return
}

func (m *User) Delete() (err error) {
	tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Model(&m).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func (m *User) DelBatch(ids []int) (err error) {
	tx := Db.Begin()
	if len(ids) > 0 {
		err = tx.Model(&m).Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func (m *User) FindById(id int) (user User, err error) {
	err = Db.Where("id=?", id).First(&user).Error
	return
}

func (m *User) FindByMaps(offset, limit int64, dataMap map[string]interface{},orderBy string) (user []User, total int, err error) {
	query := Db
	if status,isExist:=dataMap["status"].(int);isExist == true{
		query = query.Where("status = ?", status)
	}
	if nickname,ok:=dataMap["nickname"].(string);ok{
		query = query.Where("nickname LIKE ?", "%"+nickname+"%")
	}
	
	if userName,ok:=dataMap["user_name"].(string);ok{
		query = query.Where("user_name LIKE ?", "%"+userName+"%")
	}
	if startTime,ok:=dataMap["start_time"].(int64);ok{
		query = query.Where("created_at > ?", startTime)
	}
	if endTime,ok:=dataMap["end_time"].(int64);ok{
		query = query.Where("created_at <= ?", endTime)
	}
	if phone,ok:=dataMap["phone"].(string);ok{
		query = query.Where("phone LIKE ?", "%"+phone+"%")
	}
	
	if orderBy!=""{
		query = query.Order(orderBy)
	}
	
	// 获取取指page，指定pagesize的记录
	err = query.Offset(offset).Limit(limit).Find(&user).Error
	if err == nil{
		err = query.Model(&m).Count(&total).Error
	}
	return
}

/*****************************************************************新增加的方法*****************************************************************/

func (m *User) FindByUserName(user_name string) (user User, err error) {
	err = Db.Select("id,nickname,user_name,password,salt").Where("user_name=?", user_name).First(&user).Error
	return
}

//验证用户信息
func checkUser(m *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&m)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}
