package handlers

import (
	"net/http"

	"github.com/ghibbo/speedy"
)

type Handlers struct {
	App *speedy.Speedy
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}
