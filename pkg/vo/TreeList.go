package vo

type TreeList struct {
	Id int			        `json:"id"`
	MenuName string		    `json:"menu_name"`
	ParentId int			`json:"parent_id"`
	Children []*TreeList	`json:"children"`
}
