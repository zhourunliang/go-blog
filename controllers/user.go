package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) LoginPage() {
	beego.ReadFromRequest(&this.Controller)
	loginuser := this.GetSession("loginuser")
	this.Data["loginuser"] = loginuser
	this.Layout = "layout.html"
	this.TplName = "user/login.html"
}

func (this *UserController) Login() {
	flash := beego.NewFlash()
	username, password := this.Input().Get("name"), this.Input().Get("password")
	if flag, user := models.Login(username); flag && user.Password == password {
		this.SetSession("loginuser", username)
		this.Redirect("/", 302)
	} else {
	  flash.Error("用户名或密码错误")
	  flash.Store(&this.Controller)
	  this.Redirect("/login", 302)
	}
}

func (this *UserController) RegisterPage() {
	beego.ReadFromRequest(&this.Controller)
	loginuser := this.GetSession("loginuser")
	this.Data["loginuser"] = loginuser
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
		beego.Notice("用户名或密码不能为空")

		flash.Error("用户名或密码不能为空")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
		return
	} else if password != repassword {
		flash.Error("用户密码不同")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)	
		return	
	} else if flag, _ := models.FindUserByUserName(username); flag {
		flash.Error("用户名已被注册")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
	  }else {
		var user models.User
		user.Name = username
		user.Password = password
		models.SaveUser(&user)
		this.Ctx.Redirect(302, "/")
	}


}

func (this *UserController) Logout() {
	//删除指定的session   
	this.DelSession("loginuser")
	//销毁全部的session
	this.DestroySession()
	this.Redirect("/", 302)
}
