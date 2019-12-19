package vo

import "time"

type MenuItem struct {
	Id  	  int       `json:"id" form:"menu_name" gorm:"default:''"`
	MenuName  string    `json:"menu_name" form:"menu_name" gorm:"default:''"`
	ParentId  int       `json:"parent_id" form:"parent_id" gorm:"default:'0'"`
	OrderNum  int       `json:"order_num" form:"order_num" gorm:"default:'0'"`
	Url       string    `json:"url"       form:"url"       gorm:"default:'#'"`
	MenuType  int       `json:"menu_type" form:"menu_type" gorm:"default:''"`
	Visible   int       `json:"visible"   form:"visible"   gorm:"default:'0'"`
	Perms     string    `json:"perms"     form:"perms"     gorm:"default:''"`
	Icon      string    `json:"icon"      form:"icon"      gorm:"default:'#'"`
	IsFrame   int       `json:"is_frame"  form:"is_frame"  gorm:"default:'0'"`
	Component string    `json:"component" form:"component" gorm:"default:''"`
	CreateBy  int       `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt time.Time `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  int       `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt time.Time `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	RouteName string    `json:"route_name"    form:"route_name"    gorm:"default:''"`
	RoutePath string    `json:"route_path"    form:"route_path"    gorm:"default:''"`
	RouteCache     int    `json:"route_cache"       form:"route_cache"        gorm:"default:''"`
	RouteComponent string `json:"route_component"   form:"route_component"    gorm:"default:''"`
	ChildrenList []MenuItem `json:"children_list"`
}
