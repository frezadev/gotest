package controllers

import (
	. "github.com/frezadev/gotest/library/server"
	_ "github.com/lib/pq"
)

type AboutController struct {
	Config Config
}

func CreateAboutController() *AboutController {
	return new(AboutController)
}

func (c *AboutController) Default(a *AppWebContext) interface{} {
	a.Config.OutputType = OutHTML
	a.Config.OutputFile = ViewsPath + "about.html"
	a.Config.IncludeFiles = []string{ViewsPath + "_header.html", ViewsPath + "_footer.html", ViewsPath + "_menu.html"}
	return nil
}
