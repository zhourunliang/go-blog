package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) LoginPage() {
	this.Layout = "layout.html"
	this.TplName = "user/login.html"
}

func (this *UserController) Login() {

}

func (this *UserController) RegisterPage() {
	this.Layout = "layout.html"
	this.TplName = "user/register.html"
}

func (this *UserController) Register() {
	
	flash := beego.NewFlash()
	inputs := this.Input()

	username := inputs.Get("name")
	password := inputs.Get("password")
	repassword := inputs.Get("repassword")
	
	beego.Notice("username_len")
	beego.Notice(len(username))

	beego.Notice("password_len")
	beego.Notice(len(password))
	if len(username) == 0 || len(password) == 0 {
		flash.Error("用户名或密码不能为空")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
	} else if password != repassword {
		flash.Error("用户名或密码不能为空")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)		
	} else {
		var user models.User
		user.Name = username
		user.Password = password
		models.SaveUser(&user)
		this.Ctx.Redirect(302, "/")
	}


}

func (this *UserController) Logout() {

}
