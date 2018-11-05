package main

import (
    "log"
    "fmt"
    "net/http"
    "github.com/ahmeturun/goApi/persistance"
    "github.com/gorilla/mux"
)
// our main function
func main() {
    db_manager.OpenConnection()
	router := mux.NewRouter()
	router.HandleFunc("/login", Login).Methods("POST")
    router.HandleFunc("/Favourites/{userid}", GetFavourites).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func Login(w http.ResponseWriter, r *http.Request) {
    // TO-DO: Authenticate via credentials and
    //        return authentication and refresh token
    vars = mux.Vars(r)
    query = fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME=%s AND PASSWORD=%s", vars["username"], vars["password"])
    rows, err := db_manager.RunQuery(query)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, err.Error())
    }else if rows.rowsi < 1{
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintf(w, "Email or password wrong!")
    }
    else{
        http.Redirect(w, r, "homepage", httpStatusSeeOther)
    }

}
func GetFavourites(w http.ResponseWriter, r *http.Request) {
    // check tokens and return favourites
}