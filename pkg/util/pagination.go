package util

// PageData ...
type PageData struct {
	Page Page        `json:"page"`
	List interface{} `json:"list"`
}

// Page ...
type Page struct {
	PageNo     int64 `json:"page_no"`
	PageSize   int64 `json:"page_size"`
	TotalPage  int64 `json:"tatal_page"`
	TotalCount int64 `json:"tatal_count"`
	FirstPage  bool  `json:"first_page"`
	LastPage   bool  `json:"last_page"`
}

// PageUtil ...
func PageUtil(count int64, pageNo int64, pageSize int64, list interface{}) PageData {

	if pageNo <= 0 {
		pageNo = 1
	}

	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}

	page := Page{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalPage:  tp,
		TotalCount: count,
		FirstPage:  pageNo == 1,
		LastPage:   pageNo == tp,
	}

	return PageData{Page: page, List: list}
}
