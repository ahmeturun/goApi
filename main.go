package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", Login).Methods("POST")
    router.HandleFunc("/Favourites/{userid}", GetFavourites).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func Login(w http.ResponseWriter, r *http.Request) {
    // Authenticate via credentials and
    // return authentication and refresh token
}
func GetFavourites(w http.ResponseWriter, r *http.Request) {
    // check tokens and return favourites
}