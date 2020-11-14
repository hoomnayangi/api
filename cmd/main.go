package main

import (
	"net/http"

	"github.com/hoomnayangi/api/src/handler"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/phuwn/tools/db"
	"github.com/phuwn/tools/log"
)

func main() {
	db.Start()
	defer db.Close()

	addr := ":8080"
	log.Status("listening on port%s", addr)

	err := http.ListenAndServe(addr, handler.Router())
	if err != nil {
		log.Status("server got terminated, err: %s", err.Error())
	}
}
