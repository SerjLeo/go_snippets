package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

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


    app := &application {
        errorLog: errorLog,
        infoLog: infoLog,
    }

	infoLog.Printf("Server is listening on port %s", *addr)
	fileLog.Printf("Server is listening on port %s", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
