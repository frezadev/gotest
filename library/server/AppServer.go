package server

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

const (
	OutHTML = 1
	OutJSON = 2
)

type Config struct {
	OutputType   int
	OutputFile   string
	IncludeFiles []string
}

type FnContent func(r *AppWebContext) interface{}

type ServerConf struct {
	Address string
}

type AppServer struct {
	Name        string
	controllers map[string]interface{}
}

func NewApp(name string) *AppServer {
	app := new(AppServer)
	app.Name = name
	return app
}

func (a *AppServer) Register(c interface{}) error {
	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("Unable to register %v, type is not pointer \n", c)
	}

	name := strings.ToLower(reflect.Indirect(v).Type().Name())
	a.Controllers()[name] = c
	return nil
}

func (a *AppServer) Controllers() map[string]interface{} {
	if a.controllers == nil {
		a.controllers = map[string]interface{}{}
	}
	return a.controllers
}

func Start(conf *ServerConf, app *AppServer) {
	if conf != nil && app != nil {
		fmt.Println()
		log.Println("Start Hello World Server...")
		appname := strings.ToLower(app.Name)
		pattern := "/"
		if appname != "" {
			pattern = pattern + appname + "/"
		}

		for uri, handler := range app.Controllers() {
			setHandler(pattern+uri, handler)
		}

		log.Printf("listening on %v \n", conf.Address)
		http.ListenAndServe(conf.Address, nil)
	}
}

func handle(path string, fnc FnContent) {
	log.Printf("Registered Controller: %v \n", path)
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %#v\n", path)
		cfg := new(AppWebContext)
		cfg.Response = w
		cfg.Request = r

		cfg.Write(fnc(cfg))
	})
}

func setHandler(path string, handler interface{}) {
	v := reflect.ValueOf(handler)
	t := reflect.TypeOf(handler)

	methodCount := t.NumMethod()
	// log.Printf("methodCount: %v \n", methodCount)

	for mi := 0; mi < methodCount; mi++ {
		method := t.Method(mi)
		isFnContent := false

		tm := method.Type
		if tm.NumIn() == 2 && tm.In(1).String() == "*server.AppWebContext" {
			if tm.NumOut() == 1 && tm.Out(0).Kind() == reflect.Interface {
				isFnContent = true
			}
		}

		if isFnContent {
			var fnc FnContent
			fnc = v.MethodByName(method.Name).Interface().(func(*AppWebContext) interface{})

			methodName := method.Name
			handlerPath := path + "/" + strings.ToLower(methodName)

			handle(handlerPath, fnc)
		}
	}
}

func CreateResult(success bool, data interface{}, message string) map[string]interface{} {
	if !success {
		log.Println("ERROR! ", message)
	}

	return map[string]interface{}{
		"data":    data,
		"success": success,
		"message": message,
	}
}
