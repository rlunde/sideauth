package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var ServiceStartTime time.Time

func main() {
	RunService() // see authapi.go
}

//VERSION -- the version of the service
const VERSION = "0.1.0"

//APIVERSION -- the version of the sideauth API
const APIVERSION = "v1"

func init() {
	ServiceStartTime = time.Now()
}

/*RunService runs the main service endpoints
 */
func RunService() {

	// next 3 lines show use of Gorialla mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", status)
	router.HandleFunc("/version", version)

	// best practices generally recommend using only plural nouns for both GET and POST
	router.HandleFunc("/accounts", CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{account}", UpdateAccount).Methods("PUT")

	// router.HandleFunc("/", indexPage).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))

}

/*Login - used when validating pwhash(we don't need email for this)  */
type Status struct {
	Database       string `json:"database"`
	Uptime         string `json:"uptime" `
	Duration       string `json:"duration"`
	ServiceVersion string `json:"serviceVersion"`
}

func status(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// TODO: report back on the status of the database
	s := Status{}
	s.Database = "unknown"
	s.Uptime = t.Sub(ServiceStartTime).String()
	s.Duration = time.Now().Sub(t).String()
	s.ServiceVersion = VERSION
	bs, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("%s", err.Error()) //TODO: log this to a file
	}
	// TODO: return uptime, time to handle this request, service version
	io.WriteString(w, string(bs))
}
func version(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"version\": \"%s\"}", VERSION)
}
