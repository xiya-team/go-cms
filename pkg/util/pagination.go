package util

type PageData struct {
	Page Page
	List       interface{}
}

type Page struct {
	PageNo     int64
	PageSize   int64
	TotalPage  int64
	TotalCount int64
	FirstPage  bool
	LastPage   bool
}

func PageUtil(count int64, pageNo int64, pageSize int64, list interface{}) PageData {
	tp := count / pageSize
	if count % pageSize > 0 {
		tp = count / pageSize + 1
	}
	
	page := Page{}
	page.PageNo = pageNo
	page.PageSize = pageSize
	page.TotalPage = tp
	page.TotalCount = count
	page.FirstPage = pageNo == 1
	page.LastPage = pageNo == tp
	
	return PageData{Page:page, List: list}
}