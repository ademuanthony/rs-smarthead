package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "site/default.html"
}

func (c *MainController) Pricing() {
	c.TplName = "site/pricing.html"
}

func (c *MainController) Support() {
	c.TplName = "site/support.html"
}