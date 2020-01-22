package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", RootHandler)
    http.ListenAndServe(":8000", r)
}

// Temporary
func RootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}
