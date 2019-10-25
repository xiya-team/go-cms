package vo

import "time"

type DeptItem struct {
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	ParentId  int       `json:"parent_id" form:"parent_id" gorm:"default:'0'"`
	Ancestors string    `json:"ancestors" form:"ancestors" gorm:"default:''"`
	DeptName  string    `json:"dept_name" form:"dept_name" gorm:"default:''"`
	OrderNum  int       `json:"order_num" form:"order_num" gorm:"default:'0'"`
	Leader    string    `json:"leader"    form:"leader"    gorm:"default:''"`
	Phone     string    `json:"phone"     form:"phone"     gorm:"default:''"`
	Email     string    `json:"email"     form:"email"     gorm:"default:''"`
	Status    string    `json:"status"    form:"status"    gorm:"default:'0'"`
	DelFlag   string    `json:"del_flag"  form:"del_flag"  gorm:"default:'0'"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt time.Time `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt time.Time `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	ChildrenList []DeptItem `json:"children_list"`
}
