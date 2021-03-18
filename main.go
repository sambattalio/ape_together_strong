package main

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

type Workout struct {
    Exercise string `json:"exercise"`
    Weight   string `json:"weight"`
}

// i don't like this
var workouts []Workout


func main() {
    r := buildRouter()
    http.ListenAndServe(":8080", r)
}

func buildRouter() *mux.Router {
    r := mux.NewRouter()

    // set up RESTful endpoints
    r.HandleFunc("/hello", handler).Methods("GET")
    r.HandleFunc("/workout", createWorkoutHandler).Methods("POST")
    r.HandleFunc("/workout", getWorkoutHandler).Methods("GET")

    // load up static files
    staticFileDir := http.Dir("./assets/")
    staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))

    r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

    return r
}

func handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}


func getWorkoutHandler(w http.ResponseWriter, r *http.Request) {
    workoutListBytes, err := json.Marshal(workouts)

    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
	w.WriteHeader(http.StatusInternalServerError)
	return
    }

    w.Write(workoutListBytes)
}

func createWorkoutHandler(w http.ResponseWriter, r *http.Request) {
    workout := Workout{}

    err := r.ParseForm()

    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
	w.WriteHeader(http.StatusInternalServerError)
	return
    }

    workout.Exercise = r.Form.Get("exercise")
    workout.Weight   = r.Form.Get("weight")

    fmt.Println(workout.Exercise)

    workouts = append(workouts, workout)


    http.Redirect(w, r, "/assets/", http.StatusFound)
}
