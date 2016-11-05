package controllers

import (
	. "github.com/frezadev/gotest/library/server"
	_ "github.com/lib/pq"
)

type HomeController struct {
	Config Config
}

func CreateHomeController() *HomeController {
	return new(HomeController)
}

func (c *HomeController) Default(a *AppWebContext) interface{} {
	a.Config.OutputType = OutHTML
	a.Config.OutputFile = ViewsPath + "home.html"
	a.Config.IncludeFiles = []string{ViewsPath + "_header.html", ViewsPath + "_footer.html", ViewsPath + "_menu.html"}
	return nil
}
