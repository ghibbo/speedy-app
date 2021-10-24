package main

import (
	"log"
	"myapp/handlers"
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

	myHandlers := &handlers.Handlers{
		App: speed,
	}

	app := &application{
		App:      speed,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
