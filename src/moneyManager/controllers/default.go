package controllers

import (
	"github.com/astaxie/beego"
)

type Person struct {
	Name  string
	Money float64
}

type MainController struct {
	beego.Controller
}

type MyController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func (c *MyController) Get() {

	ss := []Person{{"倪春", 42}, {"倪春", 42}, {"倪春", 42}, {"倪春", 42}, {"倪春", 42}, {"倪春", 42}}

	c.Data["s"] = ss
	c.TplNames = "home.tpl"

}
