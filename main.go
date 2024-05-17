package main

import (
    "fmt"
    "html/template"
    "net/http"
    "path/filepath"
)

// PageData holds data to be passed to templates
type PageData struct {
    Title   string
    Content string
}

// renderTemplate renders the specified template with the provided data
func renderTemplate(w http.ResponseWriter, tmpl string, data *PageData) {
    tmplPath := filepath.Join("templates", tmpl+".html")
    templates, err := template.ParseFiles(tmplPath, "templates/header.html", "templates/footer.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = templates.ExecuteTemplate(w, tmpl+".html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// homeHandler handles requests to the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "Home Page",
        Content: "Welcome to the home page!",
    }
    renderTemplate(w, "index", data)
}

// aboutHandler handles requests to the about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "About Page",
        Content: "Welcome to the about page!",
    }
    renderTemplate(w, "about", data)
}

// catchAllHandler handles requests to any other pages
func catchAllHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "404 Page",
        Content: "Sorry, the page you are looking for does not exist.",
    }
    renderTemplate(w, "404", data)
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
