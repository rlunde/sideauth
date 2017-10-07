package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
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
func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to sideauth")
	fmt.Println("Endpoint Hit: indexPage")
}

//RegisterAccount -- create a new login
func RegisterAccount(w http.ResponseWriter, r *http.Request) {
	username, email, password, err := getRegistrationData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		fmt.Printf("RegisterAccount called with username %s, email %s, password %s\n", username, email, password)
	}
	//TODO: validate that account doesn't already exist
	//TODO: try to create login and save it in database
	//TODO: create a session cookie
	//TODO: return success or error message
	//TODO: on success, send email and display a verify email form
	//TODO: on error, display error message and redirect to register form
}

//LoginWithAccount -- create a new session, or return an error
func LoginWithAccount(w http.ResponseWriter, r *http.Request) {
	username, password, err := getLoginData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		fmt.Printf("LoginWithAccount called with username %s, password %s\n", username, password)
	}
	//TODO: create a session cookie

	mgr := GetMgr()
	sess, err := mgr.SessionStart(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//TODO: return success or error message
	//TODO: verify password is correct
	sess.Set("username", username)
	http.Redirect(w, r, "/", 302)
	//TODO: on error, display error message and redirect back to login form
}

//Logout -- destroy a session
func Logout(w http.ResponseWriter, r *http.Request) {
	mgr := GetMgr()
	err := mgr.SessionEnd(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	//TODO: return success or error message
	//TODO: on error, display error message and redirect back to login form
}

/*Registration - need to use BindJSON to retrieve from gin, since now posting from React as JSON struct */
type Registration struct {
	Username     string `form:"username" json:"username" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	ConfPassword string `form:"confpassword" json:"confpassword" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	Remember     bool   `form:"remember" json:"remember" `
}

func getRegistrationData(w http.ResponseWriter, r *http.Request) (username, email, password string, err error) {

	var reg Registration
	if r.Body == nil {
		http.Error(w, "Invalid request body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = checkmail.ValidateFormat(reg.Email)
	if err != nil {
		return
	}
	// err = checkmail.ValidateHost(reg.Email)
	// if err != nil {
	// 	return
	// }
	if reg.Password != reg.ConfPassword {
		err = errors.New("Password and confirm-password do not match")
	}
	username = reg.Username
	email = reg.Email
	password = reg.Password
	return
}

/*Login - need to use BindJSON to retrieve from gin, since now posting from React as JSON struct */
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Remember bool   `form:"remember" json:"remember" `
}

func getLoginData(w http.ResponseWriter, r *http.Request) (username, password string, err error) {

	var login Login
	if r.Body == nil {
		http.Error(w, "Invalid request body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err == nil {
		fmt.Printf("Got username: %s\n", login.Username)
	}
	username = login.Username
	password = login.Password
	return
}
