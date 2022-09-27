package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", home)
    http.ListenAndServe(":9110", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome to golang")
}