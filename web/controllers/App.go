package controllers

import (
	"log"
	"os"
	"path/filepath"

	. "github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var (
	AppName     string = "web"
	AppBasePath        = func(dir string, err error) string { return dir }(os.Getwd())
	ViewsPath          = filepath.Join(AppBasePath, AppName, "views") + string(os.PathSeparator)
)

// GetConnection ...
func GetConnection() (db *DB, err error) {
	// the connection info should be in the external file under /conf, not hardcoded
	db, err = Connect("postgres", "user=postgres password=password dbname=gotest sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return
}
