package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/lcabrini/npk-common"
)

func main() {
    npk.DBConnection(config)
    r := mux.NewRouter()
    //r.HandleFunc("/", RootHandler)
    npk.SetupRoutes(r)
    http.ListenAndServe(":8000", r)
}

// Temporary
func RootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}
