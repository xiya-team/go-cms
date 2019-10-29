package models

import (
	"errors"
	"github.com/syyongx/php2go"
	"github.com/wxnacy/wgo/arrays"
	"go-cms/pkg/vo"
	"time"
)

type Menu struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	MenuName  string    `json:"menu_name" form:"menu_name" gorm:"default:''"`
	ParentId  int       `json:"parent_id" form:"parent_id" gorm:"default:'0'"`
	OrderNum  int       `json:"order_num" form:"order_num" gorm:"default:'0'"`
	Url       string    `json:"url"       form:"url"       gorm:"default:'#'"`
	MenuType  int       `json:"menu_type" form:"menu_type" gorm:"default:''"`
	Visible   int       `json:"visible"   form:"visible"   gorm:"default:'0'"`
	IsFrame   int       `json:"is_frame"  form:"is_frame"  gorm:"default:'0'"`
	Perms     string    `json:"perms"     form:"perms"     gorm:"default:''"`
	Icon      string    `json:"icon"      form:"icon"      gorm:"default:'#'"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt time.Time `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt time.Time `json:"updated_at"form:"updated_at"gorm:"default:''"`
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
	query.Select("*").Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Menu{}).Count(&count)
	return
}

func (m *Menu) Create() (newAttr Menu, err error) {

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

func (m *Menu) Update() (newAttr Menu, err error) {
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

func (m *Menu) Delete() (err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		model := NewMenu()
		menu,_:=model.FindByParentId(m.Id)
		if php2go.Empty(menu) {
			err = tx.Model(&m).Delete(m).Error
		} else {
			err = errors.New("菜单下有子菜单不能删除！")
		}
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

func (m *Menu) FindByParentId(parent_id int) (menu Menu, err error) {
	err = Db.Select("*").Where("parent_id=?", parent_id).First(&menu).Error
	return
}

func (m *Menu) DelBatch(ids []int) (err error) {
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

func (m *Menu) FindById(id int) (menu Menu, err error) {
	err = Db.Select("*").Where("id=?", id).First(&menu).Error
	return
}

func (m *Menu) FindByMap(offset, limit int64, dataMap map[string]interface{},orderBy string) (res []Menu, total int64, err error) {
	query := Db
	if visible,isExist:=dataMap["visible"].(int);isExist{
		query = query.Where("visible = ?", visible)
	}
	if menuName,ok:=dataMap["menu_name"].(string);ok{
		query = query.Where("menu_name LIKE ?", "%"+menuName+"%")
	}

	if startTime,ok:=dataMap["start_time"].(string);ok{
		query = query.Where("created_at > ?", startTime)
	}
	if endTime,ok:=dataMap["end_time"].(string);ok{
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

func (m *Menu) FindAll() (res []Menu, err error) {
	query := Db
	err = query.Find(&res).Error
	return
}

func (m *Menu) FindAllByParentId(parentId int) (res []Menu, err error)   {
	query := Db
	
	query = query.Where("parent_id = ?", parentId)
	err = query.Find(&res).Error

	return
}

func (m *Menu) FindTopMenu() []*vo.TreeList {
	var menu []Menu
	query := Db
	query = query.Where("parent_id=?",0)
	_ = query.Find(&menu).Error

	treeList := []*vo.TreeList{}
	for _, v := range menu{
		node := &vo.TreeList{
			Id:v.Id,
			MenuName:v.MenuName,
			OrderNum:v.OrderNum,
			MenuType:v.MenuType,
			Visible:v.Visible,
			CreateBy:v.CreateBy,
			CreatedAt:v.CreatedAt,
			UpdateBy:v.UpdateBy,
			Icon:v.Icon,
			UpdatedAt:v.UpdatedAt,
			Perms:v.Perms,
			Remark:v.Remark,
			Url:v.Url,
			ParentId:v.ParentId,
		}
		treeList = append(treeList, node)
	}

	return treeList
}

/**
 * 生成树型数据
 */
func (m *Menu)FindMenus(pid int) []*vo.TreeList {
	query := Db
	var menu []Menu
	query = query.Where("parent_id=?",pid)
	_ = query.Find(&menu).Error
	treeList := []*vo.TreeList{}
	for _, v := range menu{
		child := v.FindMenus(v.Id)
		node := &vo.TreeList{
			Id:v.Id,
			MenuName:v.MenuName,
			OrderNum:v.OrderNum,
			MenuType:v.MenuType,
			Visible:v.Visible,
			CreateBy:v.CreateBy,
			CreatedAt:v.CreatedAt,
			UpdateBy:v.UpdateBy,
			Icon:v.Icon,
			UpdatedAt:v.UpdatedAt,
			Perms:v.Perms,
			Remark:v.Remark,
			Url:v.Url,
			ParentId:v.ParentId,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}


func (m *Menu)FindAllChildren(pid int)  []int {
	var ids []int;
	menuData,_ := m.FindAll()

	for _, menu := range menuData {
		if pid == menu.ParentId{
			ids = append(ids, menu.Id)
		} else {
			is_exist := arrays.Contains(ids, menu.ParentId)
			if is_exist != 0 {
				ids = append(ids, menu.Id)
			}
		}
	}
	return ids
}
