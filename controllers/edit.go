package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"go-blog/models"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.html"
	this.TplName = "edit.html"
}

func (this *EditController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(inputs.Get("id"))
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now()
	beego.Notice("svae...1")
	beego.Notice(blog)
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}