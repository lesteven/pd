package main

import (
    "net/http"
    "html/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
    t,_ := template.ParseFiles("./ui/html/home.page.tmpl", "./ui/html/navbar.partial.tmpl")
    t.Execute(w, nil)
}

