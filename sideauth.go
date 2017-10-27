package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	RunService() // see authapi.go
}

/*RunService runs the main service endpoints
 */
func RunService() {
	// next 3 lines show use of Gorialla mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", ping)

	router.HandleFunc("/register", RegisterAccount).Methods("POST")
	/* session related operations: login creates a session, logout destroys one */
	router.HandleFunc("/login", LoginWithAccount).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("POST")

	/* all other operations require a valid session, and validation happens as a first step */
	router.HandleFunc("/", indexPage).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))

}
