package includes

import (
    "html/template"
    "net/http"
    "path/filepath"
)

// RenderTemplate renders the specified template with the provided data
func RenderTemplate(w http.ResponseWriter, tmpl string, data *PageData) {
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
