package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/badoux/checkmail"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to sideauth")
	fmt.Println("Endpoint Hit: indexPage")
}

//CreateAccount -- create a new login
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	account, email, pwhash, err := getAccountData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("CreateAccount called with account %s, email %s, pwhash %s\n", account, email, pwhash)

	//TODO: validate that account doesn't already exist
	//TODO: try to create login and save it in database
	//TODO: return success or error message
	//TODO: on success, send email and display a verify email form
	//TODO: on error, display error message and redirect to register form
}

//LoginWithAccount -- verify pwhash for an account, or return an error
func LoginWithAccount(w http.ResponseWriter, r *http.Request) {
	account, pwhash, err := getLoginData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("LoginWithAccount called with account %s, pwhash %s\n", account, pwhash)

	// mgr := GetMgr()
	// sess, err := mgr.SessionStart(w, r)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	//TODO: return success or error message
	//TODO: verify pwhash is correct

	http.Redirect(w, r, "/", 302)
	//TODO: on error, display error message and redirect back to login form
}

/*Account -- a unique name corresponding to an identity, an authorization. For now
  the identity will be managed (such as it is) by email,
  and the authorization will be done via a password hash. Sideauth will never get the
  actual password, just the hash from the service.
*/
type Account struct {
	Account string `form:"account" json:"account" binding:"required"`
	Pwhash  string `form:"pwhash" json:"pwhash" binding:"required"`
	Email   string `form:"email" json:"email" binding:"required"`
}

func getAccountData(w http.ResponseWriter, r *http.Request) (account, email, pwhash string, err error) {

	var acct Account
	if r.Body == nil {
		http.Error(w, "Invalid request body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&acct)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = checkmail.ValidateFormat(acct.Email)
	if err != nil {
		return
	}
	// err = checkmail.ValidateHost(acct.Email)
	// if err != nil {
	// 	return
	// }
	account = acct.Account
	email = acct.Email
	pwhash = acct.Pwhash
	return
}

/*Login - used when validating pwhash(we don't need email for this)  */
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
