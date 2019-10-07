package commons

import (
	"github.com/astaxie/beego/cache"
	captchaBase "github.com/astaxie/beego/utils/captcha"
)

// 全局验证码结构体
var captcha *captchaBase.Captcha

func init()  {
	// 验证码功能
	// 使用Beego缓存存储验证码数据
	store := cache.NewMemoryCache()
	// 创建验证码
	captcha = captchaBase.NewWithFilter("/api/common/captcha", store)
	// 设置验证码长度
	captcha.ChallengeNums = 4
	// 设置验证码模板高度
	captcha.StdHeight = 50
	// 设置验证码模板宽度
	captcha.StdWidth = 120
}


