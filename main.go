package main

import (
    "log"
    "net/http"
	"encoding/json"
    "github.com/gorilla/mux"
)
type User struct {
    Name   string `json:"name"`
    Carnet string `json:"carnet"`
}
func main() {
    // Create a new Gorilla Mux router
    r := mux.NewRouter()

    // Define your routes
	r.PathPrefix("/getPics/").Handler(http.StripPrefix("/getPics/", http.FileServer(http.Dir("./images"))))
    r.HandleFunc("/", GetUserHandler).Methods("GET")

    // Start the HTTP server
    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to the home page!"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // Create a sample user
    user := User{
        Name: "Pablo García",
		Carnet: "201901107",
    }

    // Marshal user data into JSON format
    jsonData, err := json.Marshal(user)
    if err != nil {
        http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
        return
    }

    // Set content type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Send the JSON response
    w.Write(jsonData)
}