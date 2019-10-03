package util

import (
	"github.com/qinxin0720/alisms-go/SmsClient"
	"log"
	"net/http"
)

func SendAliyunSMS(){
	
	sc, err := SmsClient.NewSMSClient("oIlPLqoEBciuy86h", "NI8XDGEpBleP76CzbL3NEP1GLkTzHR")
	if err != nil {
		return
	}
	statusCode, _, _ := sc.SendSMS(SmsClient.Params{"13911052021", "洗牙网", "SMS_165117151", `{"code":"123456"}`})
	if statusCode == http.StatusOK {
		log.Println("发送成功")
	} else {
		log.Println("发送失败")
	}
}