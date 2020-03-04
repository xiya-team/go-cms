package models

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/xiya-team/helpers"
	"log"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Model
	Id          int         `json:"id"         form:"id"         gorm:"default:''"`
	Nickname    string      `json:"nickname"   form:"nickname"   gorm:"default:''" validate:"required"`
	UserName    string      `json:"user_name"  form:"user_name"  gorm:"default:''" validate:"required"`
	UserType    int         `json:"user_type"  form:"user_type"  gorm:"default:''"`
	Email       string      `json:"email"      form:"email"      gorm:"default:''" validate:"required,email"`
	Phone       string      `json:"phone"      form:"phone"      gorm:"default:''"`
	Phonenumber string      `json:"phonenumber"form:"phonenumber"gorm:"default:''"`
	Sex         int         `json:"sex"        form:"sex"        gorm:"default:'1'"`
	Avatar      string      `json:"avatar"     form:"avatar"     gorm:"default:''"`
	Password    string      `json:"password"   form:"password"   gorm:"default:''"`
	Salt        string      `json:"salt"       form:"salt"       gorm:"default:''"`
	Status      int         `json:"status"     form:"status"     gorm:"default:'1'"`
	DelFlag     int         `json:"del_flag"   form:"del_flag"   gorm:"default:'1'"`
	DeptId      int         `json:"dept_id"    form:"dept_id"    gorm:"default:''"`
	LoginIp     string      `json:"login_ip"   form:"login_ip"   gorm:"default:''"`
	LoginDate   time.Time   `json:"login_date" form:"login_date" gorm:"default:''"`
	CreateBy    int         `json:"create_by"  form:"create_by"  gorm:"default:''"`
	CreatedAt   time.Time   `json:"created_at" form:"created_at" gorm:"default:''"`
	UpdateBy    int         `json:"update_by"  form:"update_by"  gorm:"default:''"`
	UpdatedAt   time.Time   `json:"updated_at" form:"updated_at" gorm:"default:''"`
	DeletedAt   time.Time   `json:"deleted_at" form:"deleted_at" gorm:"default:''"`
	Remark      string      `json:"remark"     form:"remark"     gorm:"default:''"`
	UserPost    string      `json:"user_post"  form:"user_post"  gorm:"-"`   // 忽略这个字段
	UserRole    string      `json:"user_role"  form:"user_role"  gorm:"-"`   // 忽略这个字段
	NewPassword string      `json:"new_password" form:"new_password" gorm:"-"`   // 忽略这个字段
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

	up := NewUserPost()
	if !helpers.Empty(m.UserPost) {
		err = tx.Model(&up).Where("user_id=?", m.Id).Delete(up).Error
		if err == nil{
			s := strings.Split(m.UserPost, ",")
			for _, value := range s {
				post_id, _ := strconv.Atoi(value)
				user_post := UserPost{
					UserId:m.Id,
					PostId:post_id,
				}
				_ = tx.Model(&up).Create(user_post).Error
			}
		}
	}

	ur := NewUserRole()
	if !helpers.Empty(m.UserRole) {
		err = tx.Model(&ur).Where("role_id=?", m.Id).Delete(ur).Error
		if err == nil{
			s := strings.Split(m.UserRole, ",")
			for _, value := range s {
				role_id, _ := strconv.Atoi(value)
				user_role := UserRole{
					UserId:m.Id,
					RoleId:role_id,
				}
				_ = tx.Model(&ur).Create(user_role).Error
			}
		}
	}

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
		up := NewUserPost()
		if !helpers.Empty(m.UserPost) {
			err = tx.Model(&up).Where("user_id=?", m.Id).Delete(up).Error
			if err == nil{
				s := strings.Split(m.UserPost, ",")
				for _, value := range s {
					post_id, _ := strconv.Atoi(value)
					user_post := UserPost{
						UserId:m.Id,
						PostId:post_id,
					}
					_ = tx.Model(&up).Create(user_post).Error
				}
			}
		}

		ur := NewUserRole()
		if !helpers.Empty(m.UserRole) {
			err = tx.Model(&ur).Where("user_id=?", m.Id).Delete(ur).Error
			if err == nil{
				s := strings.Split(m.UserRole, ",")
				for _, value := range s {
					role_id, _ := strconv.Atoi(value)
					user_role := UserRole{
						UserId:m.Id,
						RoleId:role_id,
					}
					_ = tx.Model(&ur).Create(user_role).Error
				}
			}
		}

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
	err = Db.Where(&User{Id:id}).First(&user).Error

	up := NewUserPost()
	user_posts := []UserPost{}
	err = Db.Model(&up).Where(&UserPost{UserId:m.Id}).Find(&user_posts).Error
	if err == nil {
		var user_post []string
		for _, value := range user_posts {
			user_post = append(user_post,strconv.Itoa(value.PostId))
		}
		user.UserPost = strings.Join(user_post,",")
	}

	ur := NewUserRole()
	user_roles := []UserRole{}
	err = Db.Model(&ur).Where(&UserRole{UserId:m.Id}).Find(&user_roles).Error
	if err == nil {
		var user_role []string
		for _, value := range user_roles {
			user_role=append(user_role,strconv.Itoa(value.RoleId))
		}
		user.UserRole = strings.Join(user_role,",")
	}

	return
}

func (m *User) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (user []User, total int64, err error) {
	query := Db
	if status,isExist:=dataMap["status"].(int);isExist == true{
		query = query.Where("status = ?", status)
	}

	if dept_id,isExist:=dataMap["dept_id"].(int);isExist == true{
		dept := NewDept()
		dept_ids := dept.FindAllChildren(dept_id)
		dept_ids = append(dept_ids,dept_id)
		
		query = query.Where("dept_id in (?)", dept_ids)
	}

	if nickname,ok:=dataMap["nickname"].(string);ok{
		query = query.Where("nickname LIKE ?", "%"+nickname+"%")
	}
	
	if userName,ok:=dataMap["user_name"].(string);ok{
		query = query.Where("user_name LIKE ?", "%"+userName+"%")
	}
	if startTime,ok:=dataMap["start_time"].(string);ok{
		query = query.Where("created_at > ?", startTime)
	}
	if endTime,ok:=dataMap["end_time"].(string);ok{
		query = query.Where("created_at <= ?", endTime)
	}
	if phone,ok:=dataMap["phone"].(string);ok{
		query = query.Where("phone LIKE ?", "%"+phone+"%")
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
	query.Model(&User{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&user).Error
	return
}

/*****************************************************************新增加的方法*****************************************************************/

func (m *User) FindByUserName(user_name string) (user User, err error) {
	err = Db.Select("id,nickname,user_name,password,salt").Where(&User{UserName: user_name, Status: 1}).First(&user).Error
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
