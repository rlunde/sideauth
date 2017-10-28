package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/badoux/checkmail"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to sideauth")
	fmt.Println("Endpoint Hit: indexPage")
}

//RegisterAccount -- create a new login
func RegisterAccount(w http.ResponseWriter, r *http.Request) {
	account, email, pwhash, err := getRegistrationData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("RegisterAccount called with account %s, email %s, pwhash %s\n", account, email, pwhash)

	//TODO: validate that account doesn't already exist
	//TODO: try to create login and save it in database
	//TODO: create a session cookie
	//TODO: return success or error message
	//TODO: on success, send email and display a verify email form
	//TODO: on error, display error message and redirect to register form
}

//LoginWithAccount -- create a new session, or return an error
func LoginWithAccount(w http.ResponseWriter, r *http.Request) {
	account, pwhash, err := getLoginData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("LoginWithAccount called with account %s, pwhash %s\n", account, pwhash)

	//TODO: create a session cookie

	mgr := GetMgr()
	sess, err := mgr.SessionStart(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//TODO: return success or error message
	//TODO: verify pwhash is correct
	sess.Set("account", account)
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

/*Registration - used when registering an account -- we need email for recovery only */
type Registration struct {
	Account string `form:"account" json:"account" binding:"required"`
	Pwhash  string `form:"pwhash" json:"pwhash" binding:"required"`
	Email   string `form:"email" json:"email" binding:"required"`
}

func getRegistrationData(w http.ResponseWriter, r *http.Request) (account, email, pwhash string, err error) {

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
	account = reg.Account
	email = reg.Email
	pwhash = reg.Pwhash
	return
}

/*Login - used when creating a session (we don't need email for this)  */
type Login struct {
	Account string `form:"account" json:"account" binding:"required"`
	Pwhash  string `form:"pwhash" json:"pwhash" binding:"required"`
}

func getLoginData(w http.ResponseWriter, r *http.Request) (account, pwhash string, err error) {

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
		fmt.Printf("Got account: %s\n", login.Account)
	}
	account = login.Account
	pwhash = login.Pwhash
	return
}
