package main

import (
    "log"
    "time"
    "flag"
    "net/http"
)

func main() {
    port      := flag.String("p", "6060", "the port to serve on")
    directory := flag.String("d", ".", "the directory of static file to host")
    flag.Parse()

    http.Handle("/", http.FileServer(http.Dir(*directory)))
    server:= &http.Server{
        Addr: ":" + *port,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
    log.Println("Media Server Started...")
    log.Printf("Serving the directory")
    server.ListenAndServe()
}
