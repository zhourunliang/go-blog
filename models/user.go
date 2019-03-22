package models

import (
    // "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
    Password string `orm:"size(100)"`
}

func init() {
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8", 30)

    // register model
    orm.RegisterModel(new(User))

    // create table
    orm.RunSyncdb("default", false, true)
}

func Login(username string) (bool, User) {
    o := orm.NewOrm()
    var user User
    err := o.QueryTable(user).Filter("Name", username).One(&user)
    return err != orm.ErrNoRows, user
}

func FindUserByUserName(username string) (bool, User) {
    o := orm.NewOrm()
    var user User
    err := o.QueryTable(user).Filter("Name", username).One(&user)
    return err != orm.ErrNoRows, user
}

func SaveUser(user *User)  int64{
	o := orm.NewOrm()
    id, _ := o.Insert(user)
    return id
}
