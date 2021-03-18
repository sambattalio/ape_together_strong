package main

import (
    "fmt"
    "context"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Workout struct {
    Exercise string `json:"exercise"`
    Weight   string `json:"weight"`
}


func main() {
    r := buildRouter()
    http.ListenAndServe(":8080", r)
}

func GetClient() *mongo.Database {
	client, err := mongo.Connect(
        context.Background(),
        options.Client().ApplyURI("mongodb://127.0.0.1/"),
    )

    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
    }

    return client.Database("workout")
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
    coll := GetClient().Collection("exercises")
    res, err := coll.Find(context.TODO(), bson.D{})
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
    }
    var all []Workout
    res.All(context.TODO(), &all)
    workoutListBytes, err := json.Marshal(all)

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


    // insert into mongodb
    coll := GetClient().Collection("exercises")
    _, err = coll.InsertOne(context.TODO(), workout)
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
    }

    http.Redirect(w, r, "/assets/", http.StatusFound)
}
