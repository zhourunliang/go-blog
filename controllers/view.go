package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"go-blog/models"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.html"
	this.TplName = "view.html"
}