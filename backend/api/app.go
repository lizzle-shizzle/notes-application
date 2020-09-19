package api

import (
	"fmt"
	"log"
	"net/http"
)

type App struct{}

func (a *App) Run(addr string) {
	log.Printf(fmt.Sprintf("Notes API listening on %v...\n", addr[1:]))
	log.Fatalln(http.ListenAndServe(addr, nil))
}
