package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "site/dtox_index.html"
}

func (c *MainController) Pricing() {
	c.TplName = "site/pricing.html"
}

func (c *MainController) Support() {
	c.TplName = "site/support.html"
}

func (c *MainController) About() {
	c.TplName = "site/about.html"
}

func (c *MainController) Contact() {
	c.TplName = "site/contact.html"
}

func (c *MainController) GConsole() {
	c.TplName = "site/google6058e3992c01a0e3.html"
}