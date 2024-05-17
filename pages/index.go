package pages

import (
    "net/http"
    . "wendi-go/includes"
)

// HomeHandler handles requests to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "Home Page",
        Content: "Welcome to the home page!",
    }
    RenderTemplate(w, "index", data)
}
