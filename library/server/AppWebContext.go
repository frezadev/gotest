package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type AppWebContext struct {
	Config   Config
	Response http.ResponseWriter
	Request  *http.Request
}

func (r *AppWebContext) Write(data interface{}) error {
	if r.Config.OutputType == OutHTML {
		return r.WriteHTMLTemplate(data)
	}

	if r.Config.OutputType == OutJSON {
		return r.WriteJson(data)
	}

	if r.Config.OutputType != OutHTML && r.Config.OutputType != OutJSON {
		fmt.Fprint(r.Response, data)
		return nil
	}

	return nil
}

func (r *AppWebContext) WriteJson(data interface{}) error {
	w := r.Response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(r.Response).Encode(data)
}

func (r *AppWebContext) WriteHTMLTemplate(data interface{}) error {
	w := r.Response
	cfg := r.Config

	// log.Printf("%#v \n", r)
	// w.Header().Set("Content-Type", "text/html")

	if cfg.OutputFile != "" {
		viewFile := cfg.OutputFile
		bs, e := ioutil.ReadFile(viewFile)
		if e != nil {
			return e
		}
		t, e := template.New("main").Funcs(template.FuncMap{
			"BaseUrl": func() string {
				base := "/" + "web"
				// if cfg.AppName != "" {
				// 	base += strings.ToLower(cfg.AppName)
				// }
				// if base != "/" {
				// 	base += "/"
				// }
				return base
			},
		}).Parse(string(bs))

		if e != nil {
			return e
		}

		// log.Printf("%#v \n", cfg)

		for _, includeFile := range cfg.IncludeFiles {
			includes := strings.Split(includeFile, string(os.PathSeparator))
			_, e = t.New(includes[len(includes)-1]).ParseFiles(includeFile)
			if e != nil {
				return e
			}
		}

		e = t.Execute(w, data)
		if e != nil {
			return e
		}

		if e != nil {
			return e
		}
	} else {
		return fmt.Errorf("No template define for %s", strings.ToLower(r.Request.URL.String()))
	}
	return nil
}
