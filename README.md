## go-cms 一站式后端管理系统的解决方案

[![Go Report Card](https://goreportcard.com/badge/github.com/sinksmell/LanBlog)](https://goreportcard.com/report/github.com/sinksmell/LanBlog)
[![GoDoc](https://godoc.org/github.com/sinksmell/LanBlog?status.svg)](https://godoc.org/github.com/sinksmell/LanBlog)
[![Build Status](https://travis-ci.com/sinksmell/LanBlog.svg?branch=master)](https://travis-ci.com/sinksmell/LanBlog)
![Build Status](https://img.shields.io/badge/language-go-green.svg)

### 配套前端

[d2-projects/d2-admin-xiya-go-cms](https://github.com/d2-projects/d2-admin-xiya-go-cms)

### 项目介绍

技术栈 Vue.js + axios(ajax) +Beego Restful api + gorm + Mysql + Nginx

### 演示地址

[https://d2.pub/d2-admin-xiya-go-cms/preview](https://d2.pub/d2-admin-xiya-go-cms/preview)

### 安装&使用

> 以Ubuntu为例

### 简单部署

> 下载对应的 压缩包 解压运行 具体步骤待补充...
> 

### 手动编译安装

**Step1 安装mysql**

``` shell
sudo apt update
sudo apt install mysql-server mysql-common mysql-client
```

安装完成后创建数据库或者其他名称,与项目目录conf下app.conf中配置保持一致即可

``` shell
// 进入mysql后创建数据库
mysql -u root -p
// 创建完成后退出
create database go_cms;
exit;
``` 

**Step2 安装Nginx**

``` shell
sudo apt install nginx
```

**Step 3 安装编译环境**

> 若提示没有权限,请以root身份运行

* 下载并安装go语言,配置环境变量


``` shell
cd /usr/local
wget https://studygolang.com/dl/golang/go1.12.linux-amd64.tar.gz

tar zxvf  go1.12.linux-amd64.tar.gz

echo 'export GOROOT=/usr/local/go' >> ~/.bashrc 

echo 'export GOPATH=/var/www' >> ~/.bashrc 
echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bashrc 

source ~/.bashrc
``` 

* 查看是否安装成功
> 输入go version查看go版本 输入go 查看命令提示

``` shell
go version
go 
```

如果出现以下提示,则安装成功

![](https://i.loli.net/2019/03/03/5c7b8034bbdc4.png)

* 克隆项目到本地 

``` shell

cd /var/www

mkdir src

cd src

git clone https://github.com/xiya-team/go-cms

```

* 包管理和安装依赖

1、查看go版本 如使用go modules 强烈建议升级到go 1.13

	> go version
	go version go1.13.1 windows/amd64

2、设置国内模块代理 

	推荐使用 goproxy.cn  由七牛云提供(只需设置一次)

	go env -w GOPROXY=https://goproxy.cn,direct

3、运行或编译软件  自动下载相关包

	go run main.go  
  

**Step 3 安装编译环境**

* 修改Nginx配置文件

> 后台管理 

``` conf
server {
listen 8088;
server_name localhost;
charset utf-8;
access_log /var/www/go-cms.log  main;

location / {
  root /var/www/src/LanBlog/views;
  index index.html;
}

location ~ /(script|image|img|js|fonts|css)/ {
  expires 1d;
  root /var/www/src/go-cms/static/ ;
}

location /api {
  proxy_pass   http://localhost:9999/v1;
  add_header Access-Control-Allow-Methods *;
  add_header Access-Control-Max-Age 3600;
  add_header Access-Control-Allow-Credentials true;
  add_header Access-Control-Allow-Origin $http_origin;
  add_header Access-Control-Allow-Headers $http_access_control_request_headers;
  if ($request_method = OPTIONS ) {
    return 200;
  }
}   

}
```

>运行项目 

[gowatch](https://gitee.com/silenceper/gowatch)
```
go get github.com/silenceper/gowatch

gowatch //安装 gowatch 后才可以使用这个命令，不然只能使用

go run main.go // go 命令
```

### linux/mac 环境 快速运行

```
1、将项目拉到本地 git clone https://github.com/xiya-team/go-cms
2、copy conf/app-backup.conf 为 conf/app.conf 同时修改conf/app.conf中相关配置,导入/data下的数据库文件到数据库中
3、执行 ./run.sh start 即可启动项目
4、执行 ./run.sh stop 停止运行
5、执行 ./run.sh restart 重启
```

### pm2 启动方式

```
1、安装pm2
2、pm2 start go-cms (启动)
3、pm2 restart go-cms (重启)
4、pm2 stop go-cms (停止)
```

### 支持快速生成代码

```
1、需要  开启gii 设置 gii = true (运行项目时需要对gii设置为false)   
2、sudo go run main.go -m models -c controllers/sys -t menu
```

**感谢以下开源项目作者及参与者的无私奉献**

* [Beego](https://github.com/astaxie/beego/)
* [Gorm](https://github.com/jinzhu/gorm)
* [Vue](https://github.com/vuejs/vue)
* [D2Admin](https://github.com/d2-projects/d2-admin)
* 其他相关开源项目
* [Uber Go 语言代码风格指南](https://github.com/uber-go/guide)


<a href="https://github.com/d2-projects/d2-admin" target="_blank"><img src="https://raw.githubusercontent.com/FairyEver/d2-admin/master/doc/image/d2-admin@2x.png" width="200"></a>

需要进微信群 (入群验证信息: `加入go-cms群`)

<img src="public/Wechat.jpeg" width="300px"/>
