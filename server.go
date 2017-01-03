package main

import (
    "log"
    "time"
    "net/http"
    "io/ioutil"
)

func homePageHandler(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Set("Content-Type", "text/plain")
    data, err:= ioutil.ReadFile("robots.txt")
    if err != nil {
        log.Fatal(err)
    }
    rw.Write(data)
}

func main() {
    http.HandleFunc("/", homePageHandler)
    server:= &http.Server{
        Addr: ":6060",
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
    log.Println("Media Server Started...")
    server.ListenAndServe()
}
