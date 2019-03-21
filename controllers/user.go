package controllers

import (
	"github.com/astaxie/beego"
	// "go-blog/models"
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

}

func (this *UserController) Logout() {

}
