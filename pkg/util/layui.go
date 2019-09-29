package util

type LayuiTreeData struct {
	Id       int        `json:"id"`
	Title    string     `json:"title"`
	Checked  bool       `json:"checked"`
	Spread   bool       `json:"spread"`
	Children []LayuiTreeData `json:"children"`
}

type LayuiTreeDataTpl struct {
	Id   int
	Pid  int
	Name string
}

func Tree(res []LayuiTreeDataTpl, treeData *[]LayuiTreeData, pid int) (data []LayuiTreeData) {
	for _, v := range res {
		if v.Pid == pid {
			d := LayuiTreeData{Id: v.Id, Title: v.Name, Checked: true, Spread: true}
			d.Children = Tree(res, treeData, v.Id)
			data = append(data, d)
			*treeData = append(*treeData, d)
		}
	}
	return
}
