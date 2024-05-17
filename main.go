package main

import (
    "fmt"
    "net/http"
    . "wendi-go/pages"
)

func main() {
    http.HandleFunc("/", HomeHandler)        // Home page
    http.HandleFunc("/about", AboutHandler)  // About page
    http.HandleFunc("/404", CatchAllHandler) // Catch-all for other pages

    // Define the catch-all handler last to ensure it catches any unhandled routes
    http.HandleFunc("/{path}", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/404", http.StatusSeeOther)
    })

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
