package main

import (
    "net/http"
    "html/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("./ui/html/home.page.tmpl")
    if err != nil {
        http.Error(w, "Internal Server Error", 500)
        return
    }
    t.Execute(w, nil)
}

