package main

import (
    "net/http"
    "log"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hello from go app"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Only post methods allowed", 405)
        return
    }
    w.Write([]byte("Create some snippet"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

    w.Write([]byte("Show snippet"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    log.Println("Server is listening on port 4000")

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}