package util

import (
	"github.com/qinxin0720/alisms-go/SmsClient"
	"log"
	"net/http"
)

func SendAliyunSMS(){
	
	sc, err := SmsClient.NewSMSClient("", "")
	if err != nil {
		return
	}
	statusCode, _, _ := sc.SendSMS(SmsClient.Params{"", "", "", `{"code":"123456"}`})
	if statusCode == http.StatusOK {
		log.Println("发送成功")
	} else {
		log.Println("发送失败")
	}
}