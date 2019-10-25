package d3auth

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

//获取登录地址
func (e *Auth_qq) Get_Rurl(state string) string {
	return "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" + e.Conf.Appid + "&redirect_uri=" + e.Conf.Rurl + "&state=" + state
}

//获取token
func (e *Auth_qq) Get_Token(code string) (string, error) {

	str, err := HttpGet("https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=" + e.Conf.Appid + "&client_secret=" + e.Conf.Appkey + "&code=" + code + "&redirect_uri=" + e.Conf.Rurl)
	if err != nil {
		return "", err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &Auth_qq_err_res{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return "", err
		}
		return "", errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.Error_description)

	} else {
		re, _ := regexp.Compile("access_token=(.*)&expires_in")
		newres := re.FindStringSubmatch(str)
		if len(newres) >= 2 {
			return newres[1], nil
		}
		return "", nil
	}

}

//获取第三方id
func (e *Auth_qq) Get_Me(access_token string) (*Auth_qq_me, error) {

	str, err := HttpGet("https://graph.qq.com/oauth2.0/me?access_token=" + access_token)
	if err != nil {
		return nil, err
	}
	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &Auth_qq_err_res{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.Error_description)

	} else {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &Auth_qq_me{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return nil, err
		}

		return p, nil
	}

}

//获取第三方用户信息
func (e *Auth_qq) Get_User_Info(access_token string, openid string) (string, error) {

	str, err := HttpGet("https://graph.qq.com/user/get_user_info?access_token=" + access_token + "&oauth_consumer_key=" + e.Conf.Appid + "&openid=" + openid)
	if err != nil {
		return "", err
	}
	return string(str), nil

}

//构造方法
func NewAuth_qq(config *Auth_conf) *Auth_qq {
	return &Auth_qq{
		Conf: config,
	}
}
