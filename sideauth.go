package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	RunService() // see authapi.go
}

const VERSION = "0.1.0"

/*RunService runs the main service endpoints
 */
func RunService() {
	// next 3 lines show use of Gorialla mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", ping)
	router.HandleFunc("/version", version)

	// best practices generally recommend using only plural nouns for both GET and POST
	router.HandleFunc("/accounts", RegisterAccount).Methods("POST")
	router.HandleFunc("/accounts/{account}", UpdateAccount).Methods("PUT")

	/* session related operations: login creates a session, logout destroys one */
	router.HandleFunc("/sessions", LoginWithAccount).Methods("POST")
	router.HandleFunc("/sessions", Logout).Methods("DELETE")

	/* all other operations require a valid session, and validation happens as a first step */
	router.HandleFunc("/", indexPage).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))

}
func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}
func version(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Version: %s\n", VERSION)
}
