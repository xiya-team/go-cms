package SmsClient

import (
    "errors"
)

//DYSMSAPI_ENDPOINT 阿里云短信接口URL前缀
const (
    DYSMSAPI_ENDPOINT = "http://dysmsapi.aliyuncs.com"
)

//Params 短信模板参数
type Params struct {
    PhoneNumbers,
    SignName,
    TemplateCode,
    TemplateParam string
}

//SMSClient ...
type SMSClient struct {
    accessKeyID,
    secretAccessKey string
    dysmsapiClient *dysmsapiClient
}

//NewSMSClient ...
func NewSMSClient(accessKeyID, secretAccessKey string) (*SMSClient, error) {
    if accessKeyID == "" {
        return nil, errors.New("accessKeyId is empty")
    }
    if secretAccessKey == "" {
        return nil, errors.New("secretAccessKey is empty")
    }
    dsmsc, err := newDysmsapiClient(accessKeyID, secretAccessKey, DYSMSAPI_ENDPOINT)
    if err != nil {
        return nil, err
    }
    return &SMSClient{
        accessKeyID:     accessKeyID,
        secretAccessKey: secretAccessKey,
        dysmsapiClient:  dsmsc,
    }, nil
}

//SendSMS 发送短信
func (sc *SMSClient) SendSMS(params Params) (int, string, error) {
    return sc.dysmsapiClient.SendSms(params, sc.accessKeyID, sc.secretAccessKey)
}
