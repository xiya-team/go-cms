package d3auth

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

//获取登录地址
func (e *Auth_wx) Get_Rurl(state string) string {
	return "https://open.weixin.qq.com/connect/qrconnect?appid=" + e.Conf.Appid + "&redirect_uri=" + e.Conf.Rurl + "&response_type=code&scope=snsapi_login&state=" + state
}

//获取token
func (e *Auth_wx) Get_Token(code string) (*Auth_wx_succ_res, error) {

	str, err := HttpPost("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + e.Conf.Appid + "&secret=" + e.Conf.Appkey + "&code=" + code + "&grant_type=authorization_code")
	if err != nil {
		return nil, err
	}

	ismatch, _ := regexp.MatchString("errcode", str)
	if ismatch {

		p := &Auth_wx_err_res{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.Error_description)

	} else {

		p := &Auth_wx_succ_res{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}

}

//获取第三方用户信息
func (e *Auth_wx) Get_User_Info(access_token string, openid string) (string, error) {

	str, err := HttpGet("https://api.weixin.qq.com/sns/userinfo?access_token=" + access_token + "&openid=" + openid)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

//构造方法
func NewAuth_wx(config *Auth_conf) *Auth_wx {
	return &Auth_wx{
		Conf: config,
	}
}
