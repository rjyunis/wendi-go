package pages

import (
    "net/http"
    . "wendi-go/includes"
)

// CatchAllHandler handles requests to any other pages
func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "404 Page",
        Content: "Sorry, the page you are looking for does not exist.",
    }
    RenderTemplate(w, "404", data)
}
