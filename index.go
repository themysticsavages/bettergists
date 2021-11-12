package main

import (
    "./webapp"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", webapp.IndexHandler)
    log.Println("Website is up")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
