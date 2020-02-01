package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/lcabrini/npk-common"
)

func main() {
    if _, err := npk.DBConnection(config); err != nil {
        log.Printf("database: %v", err)
        os.Exit(1)
    } else {
        log.Printf("connected to database")
    }

    npk.SetupSessionStore(config)
    r := mux.NewRouter()
    //r.HandleFunc("/", RootHandler)
    npk.SetupRoutes(r)
    http.ListenAndServe(config.Listen, r)
}

// Temporary
func RootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}
