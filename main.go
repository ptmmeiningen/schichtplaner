package main

import (
	"github.com/ptmmeiningen/schichtplaner/app"
	_ "github.com/ptmmeiningen/schichtplaner/docs"
)

// @title Schichtplaner
// @version 0.1
// @description Golang backend API using Fiber and SQLite
// @contact.name Carsten Bröckert
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	// setup and run app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
