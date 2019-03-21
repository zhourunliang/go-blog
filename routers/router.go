package routers

import (
	"go-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//显示博客首页
	beego.Router("/", &controllers.IndexController{})
	//查看博客详细信息
	beego.Router("/view/:id([0-9]+)", &controllers.ViewController{})
	//新建博客博文
	beego.Router("/new", &controllers.NewController{})
	//删除博文
	beego.Router("/delete/:id([0-9]+)", &controllers.DeleteController{})
	//编辑博文
	beego.Router("/edit/:id([0-9]+)", &controllers.EditController{})


	beego.Router("/login", &controllers.UserController{}, "GET:LoginPage")
	beego.Router("/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/register", &controllers.UserController{}, "GET:RegisterPage")
	beego.Router("/register", &controllers.UserController{}, "POST:Register")
	beego.Router("/logout", &controllers.UserController{}, "GET:Logout")
}
