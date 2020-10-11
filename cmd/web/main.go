package main

import (
    "net/http"
    "log"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    fileServer := http.FileServer(http.Dir("./ui/static"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Server is listening on port 4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
