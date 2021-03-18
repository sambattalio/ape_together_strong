package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := buildRouter()
    http.ListenAndServe(":8080", r)
}

func buildRouter() *mux.Router {
    r := mux.NewRouter()

    // set up RESTful endpoints
    r.HandleFunc("/hello", handler).Methods("GET")
    r.HandleFunc("/workout", workoutPut).Methods("PUT")

    // load up static files
    staticFileDir := http.Dir("./assets/")
    staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))

    r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

    return r
}

func handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}

func workoutPut(w http.ResponseWriter, r *http.Request) {
    // TODO
}
