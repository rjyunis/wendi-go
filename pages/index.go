package pages

import (
    "net/http"
    . "wendi-go/includes"
)

// HomeHandler handles requests to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "Home Page",
    }
    RenderTemplate(w, "index", data)
}
