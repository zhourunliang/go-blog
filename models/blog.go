package models

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
	"github.com/astaxie/beego"
)

type Blog struct {
	Id      int `PK`
	Title   string
	Content string
	Created time.Time
}

func GetLink() beedb.Model {
	db, err := sql.Open("mymysql", "gotest/root/123456")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

func GetAll() (blogs []Blog) {
	db := GetLink()
	db.FindAll(&blogs)
	return
}

func GetBlog(id int) (blog Blog) {
	db := GetLink()
	db.Where("id=?", id).Find(&blog)
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	db := GetLink()
	db.Save(&blog)
	return bg
}

func DelBlog(blog Blog) {
	db := GetLink()
	// db.Delete(&blog)
	beego.Notice("del...1")
	beego.Notice(blog.Id)
	db.SetTable("blog").Where("id=?", blog.Id).DeleteRow()
	return
}