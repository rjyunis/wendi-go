package main

import (
    "fmt"
    "net/http"
    . "wendi-go/includes"
)


// homeHandler handles requests to the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "Home Page",
        Content: "Welcome to the home page!",
    }
    RenderTemplate(w, "index", data)
}

// aboutHandler handles requests to the about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "About Page",
        Content: "Welcome to the about page!",
    }
    RenderTemplate(w, "about", data)
}

// catchAllHandler handles requests to any other pages
func catchAllHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "404 Page",
        Content: "Sorry, the page you are looking for does not exist.",
    }
    RenderTemplate(w, "404", data)
}

func main() {
    http.HandleFunc("/", homeHandler)        // Home page
    http.HandleFunc("/about", aboutHandler)  // About page
    http.HandleFunc("/404", catchAllHandler) // Catch-all for other pages

    // Define the catch-all handler last to ensure it catches any unhandled routes
    http.HandleFunc("/{path}", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/404", http.StatusSeeOther)
    })

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
