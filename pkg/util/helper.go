package util

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/xiya-team/helpers"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numLetterBytes = "0123456789"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func ToFirstWordsUp(s string)  (str string){
	words := strings.Split(s, "_")
	for _,item:=range words {
		str += helpers.Ucfirst(item)
	}
	return str
}

func ToFirstWordsDown(s string)(str string) {
	var ss string
	for i, ch := range s {
		tmp := string(ch)
		if i==0{
			ss+= strings.ToLower(tmp)
		} else {
			if IsUpper(tmp){
				ss+= "_"
				ss+= strings.ToLower(tmp)
			}else {
				ss+= tmp
			}
		}
	}
	return ss
}

func IsUpper(code string) bool{
	return code == strings.ToUpper(code)
}

func JsonDecode(jsonStr string, structModel interface{}) error {
	decode := json.NewDecoder(strings.NewReader(jsonStr))
	err := decode.Decode(structModel)
	return err
}

func JsonEncode(structModel interface{}) (string, error) {
	jsonStr, err := json.Marshal(structModel)
	return string(jsonStr), err
}

func SHA256Encode(s string) string {
	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(s))
	return hex.EncodeToString(sha256.Sum(nil))
}

func Url(url string, params ...interface{}) string {
	queryString := ""
	for index, item := range params {
		if index%2 == 0 {
			queryString += item.(string) + "="
		} else {
			queryString += ToString(item) + "&"
		}
	}
	if url != "/" {
		url = strings.TrimRight(url, "/")
	}
	queryString = strings.TrimRight(queryString, "&")
	return url + "?" + queryString
}

func ToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	}
	return ""
}

// 将时间转换为人类可阅读的格式
func TimeDiffForHumans(t time.Time) string {
	unix := t.Unix()
	now := time.Now().Unix()
	b := now - unix
	if b < 0 {
		return t.Format("2006-01-01 15:04:05")
	}
	if b < 60 {
		return fmt.Sprintf("%d秒前", b)
	}
	// 单位：分钟
	if b < 3600 {
		b = b / 60
		return fmt.Sprintf("%d分钟前", b)
	}
	// 单位：小时
	b = b / 3600
	if b < 24 {
		return fmt.Sprintf("%d个小时前", b)
	}
	// 单位：天
	b = b / 24
	if b < 30 {
		return fmt.Sprintf("%d天前", b)
	}
	// 单位：月
	b = b / 30
	if b < 12 {
		return fmt.Sprintf("%d个月前", b)
	}
	// 单位：年
	b = b / 12
	if b > 3 {
		return t.Format("2006-01-01 15:04:05")
	}
	return fmt.Sprintf("%d年钱", b)
}

// 获取当前工作目录
func Pwd() string {
	return filepath.Dir(os.Args[0])
}

// 传入开始时间，计算结束时间
func ComputedHandlerSeconds(startTime int64) float64 {
	return float64(time.Now().UnixNano()-startTime) / 1e9
}

// 获取服务器IP
func GetLocalIp() string {
	addrSlice, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Get local IP addr failed!")
		return "127.0.0.1"
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

// 字符串转数组
func StrToArray(str, prefix string) interface{} {
	if str == "" {
		return nil
	}
	//strings.Join(strArr, ",")数组转字符串
	return strings.Split(str, prefix)
}

func randString(n int, LetterBytes string) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(LetterBytes) {
			b[i] = LetterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func RandString(n int) string {
	return randString(n, letterBytes)
}

func RandNumString(n int) string {
	return randString(n, numLetterBytes)
}

func ResultFilter(result []map[string]interface{},fields string) (data []map[string]interface{}){
	fields_data := strings.Split(fields, ",")

	for _, value := range result {
		var items map[string]interface{}
		for key, item := range value {
			if helpers.InArray(fields_data,key){
				items[key] = item
			}
		}
		data = append(data, items)
	}

	return
}