package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"go-blog/models"
)

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {

	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	blog := models.GetBlog(id)
	this.Data["Post"] = blog
	beego.Notice("del...1")
	beego.Notice(blog)
	models.DelBlog(blog)
	this.Ctx.Redirect(302, "/")

}