package d

//普通json格式
func ReturnJson(code int, msg string, data interface{}) (jsonData map[string]interface{}) {
	jsonData = make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg
	if data != nil {
		jsonData["data"] = data
	}
	return
}

//layui 后台返回需要的json格式
func LayuiJson(code int, msg string, data, count interface{}) (jsonData map[string]interface{}) {
	jsonData = make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg
	jsonData["count"] = count
	jsonData["data"] = data
	return
}

//bootstrap table 返回json
func TableJson(data, offset, limit, total interface{}) (jsonData map[string]interface{}) {
	jsonData = make(map[string]interface{}, 3)
	jsonData["rows"] = data
	jsonData["offset"] = offset
	jsonData["limit"] = limit
	jsonData["total"] = total
	return
}
