package vo

type MenuItem struct {
	MenuName     string 	`json:"menu_name"`
	ID           int    	`json:"id"`
	Url          string 	`json:"url"`
	Icon         string 	`json:"icon"`
	Active       string 	`json:"active"`
	ChildrenList []MenuItem `json:"children_list"`
}
