package main

import "github.com/ghibbo/speedy"

type application struct {
	App *speedy.Speedy
}

func main() {
	s := initApplication()
	s.App.ListenAndServe()
}
