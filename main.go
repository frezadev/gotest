package main

import (
	"os"

	"github.com/frezadev/gotest/library/server"
	"github.com/frezadev/gotest/web/controllers"
)

var (
	AppName  string = "web"
	basePath string = (func(dir string, err error) string { return dir }(os.Getwd()))
)

// main ...
func main() {

	serverConf := server.ServerConf{Address: "localhost:9090"}

	app := server.NewApp(AppName)
	app.Register(controllers.CreateHomeController())
	app.Register(controllers.CreateBrandAnalysisController())
	app.Register(controllers.CreateAboutController())

	// log.Printf("result:\n%#v \n", app.Controllers())

	server.Start(&serverConf, app)
}
