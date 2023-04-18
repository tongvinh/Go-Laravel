package main

import (
	"github.com/tongvinh/celeritas"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"

	cel.InfoLog.Println("Debug is set to", cel.Debug)

	app := &application{
		App: cel,
	}

	return app
}
