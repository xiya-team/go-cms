package commons

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"go-cms/controllers"
	"go-cms/pkg/e"
	"math/rand"
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
	res["path"] = "/" + path
	c.JsonResult(e.SUCCESS,"ok", res)
}

func (c *UploadController) Prepare() {
	c.EnableXSRF = false
}
