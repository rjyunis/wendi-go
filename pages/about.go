package pages

import (
    "net/http"
    . "wendi-go/includes"
)

// AboutHandler handles requests to the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "About Page",
        Content: "Welcome to the about page!",
    }
    RenderTemplate(w, "about", data)
}
