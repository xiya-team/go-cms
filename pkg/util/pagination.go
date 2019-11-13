package util

// PageData ...
type PageData struct {
	Page Page        `json:"page"`
	List interface{} `json:"list"`
}

// Page ...
type Page struct {
	PageNo     int64   `json:"page_no"`
	PageSize   int64   `json:"page_size"`
	TotalPage  int64   `json:"tatal_page"`
	TotalCount int64   `json:"tatal_count"`
	IsFirstPage  bool  `json:"is_first_page"`
	IsLastPage   bool  `json:"is_last_page"`
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
		PageNo:       pageNo,
		PageSize:     pageSize,
		TotalPage:    tp,
		TotalCount:   count,
		IsFirstPage:  pageNo == 1,
		IsLastPage:   pageNo == tp,
	}

	return PageData{Page: page, List: list}
}

func Pages(count int64, pageNo int64, pageSize int64) Page {

	if pageNo <= 0 {
		pageNo = 1
	}

	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}

	page := Page{
		PageNo:       pageNo,
		PageSize:     pageSize,
		TotalPage:    tp,
		TotalCount:   count,
		IsFirstPage:  pageNo == 1,
		IsLastPage:   pageNo == tp,
	}

	return page
}