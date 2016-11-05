package controllers

import (
	model "github.com/frezadev/gotest/library/models"
	. "github.com/frezadev/gotest/library/server"
	_ "github.com/lib/pq"
)

type BrandAnalysisController struct {
	Config Config
}

func CreateBrandAnalysisController() *BrandAnalysisController {
	return new(BrandAnalysisController)
}

func (c *BrandAnalysisController) Default(a *AppWebContext) interface{} {
	a.Config.OutputType = OutHTML
	a.Config.OutputFile = ViewsPath + "brand-analysis.html"
	a.Config.IncludeFiles = []string{ViewsPath + "_header.html", ViewsPath + "_footer.html", ViewsPath + "_menu.html"}
	return nil
}

func (c *BrandAnalysisController) Get(a *AppWebContext) interface{} {
	a.Config.OutputType = OutJSON
	db, err := GetConnection()

	if err != nil {
		return CreateResult(false, nil, err.Error())
	}

	bAnalysis := []model.BrandAnalysis{}
	db.Select(&bAnalysis, "SELECT * from brand_analysis ORDER BY total_user ASC")

	return CreateResult(true, bAnalysis, "success")
}
