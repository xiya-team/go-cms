package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	Id            int           `json:"id"            form:"id"            gorm:"default:''"`
	CategoryId    int           `json:"category_id"   form:"category_id"   gorm:"default:'0'"`
	PostTitle     string        `json:"post_title"    form:"post_title"    gorm:"default:''"`
	Author        string        `json:"author"        form:"author"        gorm:"default:''"`
	PostStatus    int           `json:"post_status"   form:"post_status"   gorm:"default:'1'"`
	CommentStatus int           `json:"comment_status"form:"comment_status"gorm:"default:'1'"`
	Flag          int           `json:"flag"          form:"flag"          gorm:"default:'0'"`
	PostHits      int           `json:"post_hits"     form:"post_hits"     gorm:"default:'0'"`
	PostFavorites int           `json:"post_favorites"form:"post_favorites"gorm:"default:'0'"`
	PostLike      int           `json:"post_like"     form:"post_like"     gorm:"default:'0'"`
	CommentCount  int           `json:"comment_count" form:"comment_count" gorm:"default:'0'"`
	PostKeywords  string        `json:"post_keywords" form:"post_keywords" gorm:"default:''"`
	PostExcerpt   string        `json:"post_excerpt"  form:"post_excerpt"  gorm:"default:''"`
	PostSource    string        `json:"post_source"   form:"post_source"   gorm:"default:''"`
	Image         string        `json:"image"         form:"image"         gorm:"default:''"`
	PostContent   string        `json:"post_content"  form:"post_content"  gorm:"default:''"`
	Status        int           `json:"status"        form:"status"        gorm:"default:'1'"`
	CreatedAt     time.Time     `json:"created_at"    form:"created_at"    gorm:"default:'0'"`
	UpdatedAt     time.Time     `json:"updated_at"    form:"updated_at"    gorm:"default:'0'"`
	
}


func NewArticle() (article *Article) {
	return &Article{}
}

func (m *Article) Pagination(offset, limit int, key string) (res []Article, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Article{}).Count(&count)
	return
}

func (m *Article) Create() (newAttr Article, err error) {

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

func (m *Article) Update() (newAttr Article, err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Model(&m).Where("id=?", m.Id).Save(m).Error
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

func (m *Article) Delete() (err error) {
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

func (m *Article) DelBatch(ids []int) (err error) {
    tx := Db.Begin()
	if len(ids) > 0 {
		err = tx.Model(&m).Where("id=?", m.Id).Updates(m).Error
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

func (m *Article) FindById(id int) (article Article, err error) {
	err = Db.Where("id=?", id).First(&article).Error
	return
}

func (db *Article) TakeById(id int) (res Article, err error) {
	err = Db.Where("id = ?", id).Take(&res).Error
	return
}

func (m *Article) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Article, total int64, err error) {
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
	query.Model(&Article{}).Count(&total)
	err = query.Offset(offset).Limit(limit).Find(&res).Error
	return
}

/**
 * 增加访问数
 */
func (db *Article) AllVisitCount(id int) error {
	return Db.Model(&Article{}).Where("id = ?", id).UpdateColumn("visit", gorm.Expr("visit + 1")).Error
}