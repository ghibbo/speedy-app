package main

import (
	"log"
	"os"

	"github.com/ghibbo/speedy"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init speedy
	speed := &speedy.Speedy{}
	err = speed.New(path)
	if err != nil {
		log.Fatal(err)
	}

	speed.AppName = "myapp"
	speed.Debug = true

	app := &application{
		App: speed,
	}

	return app
}
