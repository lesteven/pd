package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    port := ":3000"
    log.Println("Starting server on port", port)

    r := mux.NewRouter()

    // handle static files
    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static")))
    r.PathPrefix("/static/").Handler(fileServer)

    r.HandleFunc("/", Home)

    log.Fatal(http.ListenAndServe(port, r))
}
