package d3auth

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

//获取登录地址
func (e *Auth_wb) Get_Rurl(state string) string {
	return "https://api.weibo.com/oauth2/authorize?client_id=" + e.Conf.Appid + "&response_type=code&display=page&redirect_uri=" + e.Conf.Rurl + "&state=" + state
}

//获取token
func (e *Auth_wb) Get_Token(code string) (*Auth_wb_succ_res, error) {

	str, err := HttpPost("https://api.weibo.com/oauth2/access_token?client_id=" + e.Conf.Appid + "&client_secret=" + e.Conf.Appkey + "&code=" + code + "&grant_type=authorization_code&redirect_uri=" + e.Conf.Rurl)
	if err != nil {
		return nil, err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {

		p := &Auth_wb_err_res{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.Error_description)

	} else {

		p := &Auth_wb_succ_res{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}

}

//获取第三方用户信息
func (e *Auth_wb) Get_User_Info(access_token string, openid string) (string, error) {

	str, err := HttpGet("https://api.weibo.com/2/users/show.json?access_token=" + access_token + "&uid=" + openid)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

//构造方法
func NewAuth_wb(config *Auth_conf) *Auth_wb {
	return &Auth_wb{
		Conf: config,
	}
}
