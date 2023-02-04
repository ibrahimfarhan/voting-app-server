package main

import (
	"github.com/ibrahimfarhan/voting-app/voting-app-server/api"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/app"
)

func main() {
	initApp()
}

func initApp() {
	a := app.NewApp()
	api.InitApi(a)
	a.InitServer()
}
