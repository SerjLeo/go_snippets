package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

    flag.Parse()
    
    //Custom logs
	infoLog := log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
    
    //Logging to file
    f, err := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        errorLog.Fatal(err)
    }
    fileLog := log.New(f, "File\t", log.Ldate|log.Ltime)
    defer f.Close()

    //Create Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Server is listening on port %s", *addr)
	fileLog.Printf("Server is listening on port %s", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
