package main

import (
	"myapp/handlers"

	"github.com/ghibbo/speedy"
)

type application struct {
	App      *speedy.Speedy
	Handlers *handlers.Handlers
}

func main() {
	s := initApplication()
	s.App.ListenAndServe()
}
