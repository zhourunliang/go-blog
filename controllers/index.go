package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	loginuser := this.GetSession("loginuser")
	this.Data["loginuser"] = loginuser
	this.Data["blogs"] = models.GetAll()
	this.Layout = "layout.html"
	this.TplName = "index.html"
}