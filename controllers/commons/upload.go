package commons

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/sirupsen/logrus"
	"go-cms/controllers"
	"go-cms/pkg/e"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

type UploadController struct {
	controllers.BaseController
}

// @router /member/upload/image [post]
func (c *UploadController) Image() {
	file, header, err := c.GetFile("file")
	if err != nil {
		c.JsonResult(e.ERROR, "请选择头像文件")
	}
	defer file.Close()
	// 文件mime判断
	mime := header.Header["Content-Type"][0]
	if mime != "image/jpeg" && mime != "image/png" && mime != "image/gif" {
		c.JsonResult(e.ERROR, "请上传有效图片文件")
	}
	// 文件后缀判断
	extensions := strings.Split(header.Filename, ".")
	extension := strings.ToLower(extensions[len(extensions)-1])
	if extension != "jpg" && extension != "png" && extension != "gif" {
		c.JsonResult(e.ERROR, "请上传jpeg/png/gif图片")
	}
	// 文件大小判断
	if header.Size/(1024*1024) > 2 {
		c.JsonResult(e.ERROR, "头像文件大小不能超过2MB")
	}
	// 保存文件
	rand.Seed(time.Now().Unix())
	newFilename := fmt.Sprintf("%d+%d", time.Now().Unix(), rand.Intn(100))
	newFilename = fmt.Sprintf("%x", md5.Sum([]byte(newFilename)))
	path := path.Join("static/uploads/avatar", newFilename+"."+extension)
	err = c.SaveToFile("file", path)
	if err != nil {
		logs.Info(err)
		c.JsonResult(e.ERROR, "头像上传失败")
	}
	res := make(map[string]string)
	res["path"] = beego.AppConfig.String("app_url")+"/" + path
	c.JsonResult(e.SUCCESS,"ok", res)
}

func (c *UploadController) Prepare() {
	c.EnableXSRF = false
}

func (c *UploadController) BaiduOSS()  {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := "a12ba926fe4748ea88f8ec575dd7fece", "6c84d9e8a21f4f5fb93c35655baa5347"

	// BOS服务的Endpoint
	ENDPOINT := "http://bj.bcebos.com"

	// 创建BOS服务的Client
	bosClient, _ := bos.NewClient(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ENDPOINT)

	// 创建Bucket
	is_Exist,err := bosClient.DoesBucketExist("xiya");
	if is_Exist ==false{
		if loc, err := bosClient.PutBucket("xiya"); err != nil {
			fmt.Println("create bucket failed:", err)
		} else {
			fmt.Println("create bucket success at location:", loc)
		}
	}

	file, header, err := c.GetFile("file")
	if err != nil {
		c.JsonResult(e.ERROR, "请选择头像文件")
	}
	defer file.Close()

	// 文件mime判断
	mime := header.Header["Content-Type"][0]
	if mime != "image/jpeg" && mime != "image/png" && mime != "image/gif" {
		c.JsonResult(e.ERROR, "请上传有效图片文件")
	}

	extensions := strings.Split(header.Filename, ".")
	extension := strings.ToLower(extensions[len(extensions)-1])
	if extension != "jpg" && extension != "png" && extension != "gif" {
		c.JsonResult(e.ERROR, "请上传jpeg/png/gif图片")
	}
	// 文件大小判断
	if header.Size/(1024*1024) > 2 {
		c.JsonResult(e.ERROR, "头像文件大小不能超过2MB")
	}
	// 保存文件
	rand.Seed(time.Now().Unix())
	newFilename := fmt.Sprintf("%d+%d", time.Now().Unix(), rand.Intn(100))
	newFilename = fmt.Sprintf("%x", md5.Sum([]byte(newFilename)))

	path := path.Join("static/uploads/avatar", newFilename+"."+extension)
	err = c.SaveToFile("file", path)


	// 使用基本接口，提供必需参数从数据流上传
	bodyStream, err := bce.NewBodyFromFile(path)
	etag, err := bosClient.BasicPutObject("xiya", newFilename+"."+extension, bodyStream)

	// 上传对象
	if err != nil {
		fmt.Println("upload file to BOS failed:", err)
	}

	//删除文件
	err =os.Remove(path)
	if err!=nil{
		fmt.Println(err)
	}

	image := "http://image.xiya.vip/"+newFilename+"."+extension
	fmt.Println("upload file to BOS success, etag = ", etag)
	logrus.Debug("Useful debugging information.")
	c.JsonResult(e.SUCCESS,"ok", image)
}