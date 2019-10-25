package d3auth

//	"encoding/json"

//基本配置
type Auth_conf struct {
	Appid  string
	Appkey string
	Rurl   string
}

//@ qq 结构 ------------------------------------------------- start

type Auth_qq struct {
	Conf *Auth_conf
}
type Auth_qq_err_res struct {
	Error             int    `json:"error"`
	Error_description string `json:"error_description"`
}
type Auth_qq_me struct {
	Client_ID string `json:"client_id"`
	OpenID    string `json:"openid"`
}

//@ qq 结构 ------------------------------------------------- end

//@ weibo 结构 ------------------------------------------------- start

type Auth_wb struct {
	Conf *Auth_conf
}

type Auth_wb_err_res struct {
	Error             int    `json:"error_code"`
	Error_description string `json:"error"`
}

type Auth_wb_succ_res struct {
	Access_Token string `json:"access_token"`
	Openid       string `json:"uid"`
}

//@ weibo 结构 ------------------------------------------------- end

//@ weixin 结构 ------------------------------------------------- start

type Auth_wx struct {
	Conf *Auth_conf
}

type Auth_wx_err_res struct {
	Error             int    `json:"errcode"`
	Error_description string `json:"errmsg"`
}

type Auth_wx_succ_res struct {
	Access_Token string `json:"access_token"`
	Openid       string `json:"openid"`
}

//@ weixin 结构 ------------------------------------------------- end

//@ github 结构 ------------------------------------------------- start

type Auth_github struct {
	Conf *Auth_conf
}

type Auth_github_err_res struct {
	Error             int    `json:"errcode"`
	Error_description string `json:"errmsg"`
}

type Auth_github_succ_res struct {
	Access_Token string `json:"access_token"`
	Openid       string `json:"openid"`
}

//@ github 结构 ------------------------------------------------- end

//@ gitee 结构 ------------------------------------------------- start

type Auth_gitee struct {
	Conf *Auth_conf
}

type Auth_gitee_err_res struct {
	Error             int    `json:"errcode"`
	Error_description string `json:"errmsg"`
}

type Auth_gitee_succ_res struct {
	Access_Token string `json:"access_token"`
}

//@ gitee 结构 ------------------------------------------------- end
