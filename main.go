package main

import (
    "fmt"
    "context"
    "time"
    "strings"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/dgrijalva/jwt-go"
)

type Workout struct {
    Username string `json:"username"`
    Exercise string `json:"exercise"`
    Weight   string `json:"weight"`
    Date     int    `json:"date"`
}

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type jwtClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

type JWTPayload struct {
    Jwt string `json:"jwttoken"`
}


func main() {
    r := buildRouter()
    http.ListenAndServe(":8000", r)
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
    r.HandleFunc("/login_attempt", loginAttemptHandler).Methods("POST")
    r.HandleFunc("/check_token", tokenCheckHandler).Methods("POST")

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

    reqToken := r.Header.Get("Authorization")
    splitToken := strings.Split(reqToken, "Bearer ")
    reqToken = splitToken[1]
    // TODO query based on username

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

    // get jwt
    for name, values := range r.Header {
    // Loop over all values for the name.
    for _, value := range values {
        fmt.Println(name, value)
    }
	}
    reqToken := r.Header.Get("Authorization")
    fmt.Println(reqToken);
    //splitToken := strings.Split(reqToken, "Bearer ")
    //reqToken = splitToken[1]

    // load json data
    err := json.NewDecoder(r.Body).Decode(&workout);
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
	w.WriteHeader(http.StatusInternalServerError)
	return
    }
    fmt.Println(reqToken)
    fmt.Println(workout.Exercise);

    // insert into mongodb
    coll := GetClient().Collection("exercises")
    _, err = coll.InsertOne(context.TODO(), workout)
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
    }

    http.Redirect(w, r, "/assets/", http.StatusFound)
}

func loginAttemptHandler(w http.ResponseWriter, r *http.Request) {
    user := User{}

    err := json.NewDecoder(r.Body).Decode(&user);
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
	w.WriteHeader(http.StatusInternalServerError)
	return
    }
    // check if valid mongodb
    coll := GetClient().Collection("users")
    var result User;
    err = coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: user.Username}}).Decode(&result)
    // verify password really probably definitely not secure way
    // but hey who is actually using this
    if err != nil || user.Password != result.Password {
        fmt.Println(fmt.Errorf("account Error"))
	w.Write([]byte(""))
	return;
    }

    token, expirationTime := generateJWT(user.Username);
    http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
    })
    w.Write([]byte(fmt.Sprintf("%s", token)))
}

func tokenCheckHandler(w http.ResponseWriter, r *http.Request) {
    jwtPayload := JWTPayload{}

    err := json.NewDecoder(r.Body).Decode(&jwtPayload);
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
	w.WriteHeader(http.StatusInternalServerError)
	return
    }

    res := validateJWT(jwtPayload.Jwt)
    w.Write([]byte(fmt.Sprintf("%s", res)))
}

func generateJWT(u string) (string, time.Time) {
    // thx https://qvault.io/cryptography/how-to-build-jwts-in-go-golang/
    expirationTime := time.Now().Add(5 * time.Hour)
    claims := &jwtClaims {
        Username: u,
	StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
	},
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    signedToken, err := token.SignedString([]byte("TEMPUSEREALKEY"))

    if err != nil {
        fmt.Println(fmt.Errorf("JWT Error: %v", err))
    }

    return signedToken, expirationTime
}

func validateJWT(userJwt string) string {
    token, err := jwt.ParseWithClaims(
	userJwt,
	&jwtClaims{},
	func(token *jwt.Token) (interface{}, error) {
		return []byte("TEMPUSEREALKEY"), nil
	},
    )
    if err != nil {
        fmt.Println(fmt.Errorf("error w/ claim Error: %v", err))
	return ""
    }
    claims, ok := token.Claims.(*jwtClaims)
    if !ok {
        fmt.Println(fmt.Errorf("claims parsing error: %v", err))
        return ""
    }

    if claims.ExpiresAt < time.Now().UTC().Unix() {
        fmt.Println(fmt.Errorf("JWT expired Error: %v", err))
        return ""
    }

    username := claims.Username
    return username
}
