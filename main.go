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

func indexHandler(w http.ResponseWriter, r *http.Request) {
    data := &PageData{
        Title:   "Home Page",
        Content: "Welcome to the home page!",
    }
    renderTemplate(w, "index", data)
}

func main() {
    http.HandleFunc("/", indexHandler)
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
