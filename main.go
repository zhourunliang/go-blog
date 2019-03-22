package main

import (
	_ "go-blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

