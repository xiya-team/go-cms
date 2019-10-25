# d3auth

第三方登录golang版本的简单实现


# 安装

`go get gitee.com/zchunshan/d3auth`

[![GoDoc](http://www.godoc.org/gitee.com/zchunshan/d3auth?status.svg)](http://www.godoc.org/gitee.com/zchunshan/d3auth)

# emmm

```golang
		type Auth_conf struct {
			Appid  string
			Appkey string
			Rurl   string
		}
```

# 使用

## gitee

```golang

	giteeconf := &d3auth.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://www.oschina.net/login"}

	giteeAuth := d3auth.NewAuth_gitee(giteeconf)

	fmt.Print(giteeAuth.Get_Rurl("state")) //获取第三方登录地址

	token, err := giteeAuth.Get_Token("3ce778954a27dd95e46e999667353ac4b780d50fa987d8646a534ac820b02c13") //回调页收的code 获取token
	fmt.Print(token, err)                                                                                 
	userinfo, _ := giteeAuth.Get_User_Info(token.Access_Token)                                            //获取用户信息 userinfo 是一个json字符串返回
	fmt.Print(userinfo)


```

## qq

```golang
		qqconf := &d3auth.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://www.change.tm"}

		qqAuth := d3auth.NewAuth_qq(qqconf)

		fmt.Print(qqAuth.Get_Rurl("state")) //获取第三方登录地址

		token, err := qqAuth.Get_Token("code")  //回调页收的code 获取token


		me, err := qqAuth.Get_Me(token)  //获取第三方id

		Client_ID string `json:"client_id"`
		OpenID    string `json:"openid"`

		userinfo, _ := wbAuth.Get_User_Info(token, me.OpenID)  //获取用户信息 userinfo 是一个json字符串返回
```

## weibo

```golang
		wbconf := &d3auth.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://www.change.tm"}

		wbAuth := d3auth.NewAuth_wb(wbconf)

		fmt.Print(wbAuth.Get_Rurl("state")) //获取第三方登录地址


		tokenobj, err := wbAuth.Get_Token("code")

		Access_Token string `json:"access_token"`
		Openid       string `json:"uid"`

		userinfo, _ := wbAuth.Get_User_Info(tokenobj.Access_Token, tokenobj.Openid)//获取用户信息 userinfo 是一个json字符串返回
```

## wechat

```golang
		wxconf := &d3auth.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://www.change.tm"}

		wxAuth := d3auth.NewAuth_wx(wxconf)

		fmt.Print(wxAuth.Get_Rurl("sate")） //获取第三方登录地址

		wxres, err := wxAuth.Get_Token("code")

		userinfo, _ := wxAuth.Get_User_Info(wxres.Access_Token, wxres.Openid) //获取用户信息 userinfo 是一个json字符串返回
```

## github

```golang
		githubconf := &d3auth.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://www.change.tm/D3/d3_code/type/github"}

		githubAuth := d3auth.NewAuth_github(githubconf)

		fmt.Print(githubAuth.Get_Rurl("state"), "\r\n") //获取第三方登录地址

		token, err := githubAuth.Get_Token("6d92d879d8b1a86922a9") //回调页收的code 获取token
		if err != nil {
			fmt.Print(err, "\r\n")
			return
		}

		userinfo, err := githubAuth.Get_User_Info(token) //获取用户信息 userinfo 是一个json字符串返回
		if err != nil {
			fmt.Print(err)
		}

		fmt.Print(userinfo)	
```


