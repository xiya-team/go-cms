package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)
var location, _ = time.LoadLocation("Local")
/**
 * string转换int
 * @method parseInt
 * @param  {[type]} b string        [description]
 * @return {[type]}   [description]
 */
func ParseInt(b string, defInt int) int {
	id, err := strconv.Atoi(b)
	if err != nil {
		return defInt
	} else {
		return id
	}
}

/**
 * int转换string
 * @method parseInt
 * @param  {[type]} b string        [description]
 * @return {[type]}   [description]
 */
func ParseString(b int) string {
	id := strconv.Itoa(b)
	return id
}

/**
 * 转换浮点数为string
 * @method func
 * @param  {[type]} t *             Tools [description]
 * @return {[type]}   [description]
 */
func ParseFlostToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 5, 64)
}

/**
 * md5 加密
 * @method MD5
 * @param  {[type]} data string [description]
 */
func MD5(data string) string {
	_m := md5.New()
	io.WriteString(_m, data)
	return fmt.Sprintf("%x", _m.Sum(nil))
}

/**
 * sha1 加密
 * @method MD5
 * @param  {[type]} data string [description]
 */
// func SHA1(data string) string {
// 	_m := sha1.New()
// 	io.WriteString(_m, data)
// 	return fmt.Sprintf("%x", _m.Sum(nil))
// }

/**
 * 结构体转成json 字符串
 * @method StruckToString
 * @param  {[type]}       data interface{} [description]
 */
func StructToString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}

/**
 * 结构体转换成map对象
 * @method func
 * @param  {[type]} t *Tools        [description]
 * @return {[type]}   [description]
 */
func StructToMap(obj interface{}) map[string]interface{} {
	k := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < k.NumField(); i++ {
		data[strings.ToLower(k.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

//生成随机字符串
func GetRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()+[]{}/<>;:=.,?"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

/**
 * 字符串截取
 * @method func
 * @param  {[type]} t *Tools        [description]
 * @return {[type]}   [description]
 */
func SubString(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}

/**
 * base64 解码
 * @method func
 * @param  {[type]} t *Tools        [description]
 * @return {[type]}   [description]
 */
func Base64Decode(str string) string {
	s, err :=  base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(s)
}

/**
 * 控制台打印测试
 * @method log
 * @param  {[type]} s string        [description]
 * @return {[type]}   [description]
 */
func Logs(s string) {
	log.Printf(s)
}


func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	//随机种子 (如果不以时间戳作为时间种子, 可能每次生成的随机数每次都相同)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			// random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func GetRootPath() string {
	pwd, _ := os.Getwd()
	return pwd
}

//GetMd5 md5加密字符串
func GetMd5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}

//GetTimeStamp 获取时间戳
func GetTimeStamp() int {
	timestamp := time.Now().In(location).Unix()
	return int(timestamp)
}

//NowDate 当前时间 Y m d H:i:s
func NowDate(str string) string {
	t := time.Now().In(location)
	str = strings.Replace(str, "Y", "2006", 1)
	str = strings.Replace(str, "m", "01", 1)
	str = strings.Replace(str, "d", "02", 1)
	str = strings.Replace(str, "H", "13", 1)
	str = strings.Replace(str, "i", "04", 1)
	str = strings.Replace(str, "s", "05", 1)
	return t.Format(str)
}

//FormatTime 时间戳格式化时间
func FormatTime(timestamp int64) string {
	t := time.Unix(timestamp, 0).In(location)
	str := "Y-m-d H:i:s"
	str = strings.Replace(str, "Y", "2006", 1)
	str = strings.Replace(str, "m", "01", 1)
	str = strings.Replace(str, "d", "02", 1)
	str = strings.Replace(str, "H", "13", 1)
	str = strings.Replace(str, "i", "04", 1)
	str = strings.Replace(str, "s", "05", 1)
	return t.Format(str)
}

//GetImg 根据图片路径生成图片,待优化函数
func GetImg(path, waterPath string) {
	fmt.Println("开始处理图片")
	f, err := os.Open(path) //打开文件
	if err != nil {
		fmt.Println("打开文件失败:", err.Error())
		return
	}
	defer f.Close()
	filename := strings.Split(f.Name(), ".")
	if len(filename) != 2 || (filename[1] != "jpg" && filename[1] != "jpeg" && filename[1] != "gif" && filename[1] != "png") {
		return
	}
	fmt.Println("解析文件信息：", filename)
	
	var imager image.Image
	if filename[1] == "jpg" {
		imager, err = jpeg.Decode(f)
	} else if filename[1] == "jpeg" {
		imager, err = jpeg.Decode(f)
	} else if filename[1] == "gif" {
		imager, err = gif.Decode(f)
	} else if filename[1] == "png" {
		imager, err = png.Decode(f)
	}
	if err != nil {
		fmt.Println("打开文件失败:", err.Error())
		return
	}
	
	fmt.Println("解码文件:", imager)
	
	//获取图片缩略图
	thumbnail := resize.Thumbnail(120, 120, imager, resize.Lanczos3)
	fileThumb, err := os.Create(filename[0] + strconv.Itoa(int(time.Now().Unix())) + "_thmub.jpg")
	if err == nil {
		jpeg.Encode(fileThumb, thumbnail, &jpeg.Options{
			Quality: 80})
		fileThumb.Close()
		fmt.Println("生成缩略图片成功")
	}
	rectangler := imager.Bounds()
	fmt.Println("获取图片的0点和尾点:", rectangler)
	//创建画布
	newWidth := 200
	m := image.NewRGBA(image.Rect(0, 0, newWidth, newWidth*rectangler.Dy()/rectangler.Dx()))
	//在画布上绘制图片 m画布 m.bounds画布参数, imager 要参照打开的图片信息 image.Point 图片绘制的其实地址 绘制资源
	draw.Draw(m, m.Bounds(), imager, image.Point{100, 100}, draw.Src)
	//绘制水印图
	
	//必须是PNG图片
	
	warter, wterr := os.Open(waterPath)
	if wterr == nil {
		//无错误的时候解码
		watermark, dewaerr := png.Decode(warter)
		if dewaerr == nil {
			//无错误的时候添加水印
			draw.Draw(m, watermark.Bounds().Add(image.Pt(30, 30)), watermark, image.ZP, draw.Over)
		} else {
			fmt.Println("水印图片解码失败")
		}
	} else {
		fmt.Println("水印图片打开失败")
	}
	toimg, _ := os.Create(filename[0] + strconv.Itoa(GetTimeStamp()) + "-120-80.jpg") //创建文件系统
	defer toimg.Close()
	//toimg 保存的名称 要参照的画布，图片选项。默认透明图
	jpeg.Encode(toimg, m, &jpeg.Options{
		Quality: jpeg.DefaultQuality}) //保存为jpeg图片
}

//Password 生成密码
func Password(password, encrypt string) string {
	return GetMd5(GetMd5(password) + encrypt)
}

//IsFalse 检测字段是否为 空 0 nil
func IsFalse(args ...interface{}) bool {
	for _, v := range args {
		switch v.(type) {
		case string:
			if v == nil || v == "" {
				return false
			}
		case int, int64, int8, int32:
			if v == nil || v == 0 {
				return false
			}
		default:
			if v == nil {
				return false
			}
			
		}
	}
	return false
}

//IsError 检测是否有Error
func IsError(args ...error) bool {
	for _, v := range args {
		if v != nil {
			return true
		}
	}
	return false
}

//Struct2Map 结构体转map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj) //得到变量的类型
	v := reflect.ValueOf(obj)
	
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); /*取得字段长度*/ i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func InitPageCount(page, pageCount int64) (int64, int64) {
	if page <= 0 {
		page = 0
	}
	if pageCount <= 0 {
		pageCount = 10
	}
	return page, pageCount
}