package util

import (
	"github.com/astaxie/beego"
	"github.com/qinxin0720/alisms-go/SmsClient"
	"log"
	"net/http"
)

//TemplateCode `{"code":"123456"}`
func SendAliyunSMS(PhoneNumbers ,SignName ,TemplateParam string){
	sc, err := SmsClient.NewSMSClient(beego.AppConfig.String("sms::accessKeyID"), beego.AppConfig.String("sms::secretAccessKey"))
	if err != nil {
		return
	}
	statusCode, _, _ := sc.SendSMS(SmsClient.Params{PhoneNumbers, SignName, beego.AppConfig.String("sms::accessKeyID"), TemplateParam})
	if statusCode == http.StatusOK {
		log.Println("发送成功")
	} else {
		log.Println("发送失败")
	}
}