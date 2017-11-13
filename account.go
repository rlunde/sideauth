package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*Account -- a unique name corresponding to an identity, an authorization, and that
  can have sessions. For now the identity will be managed (such as it is) by email,
  and the authorization will be done via a password hash. Sideauth will never get the
  actual password, just the hash from the service.
*/
type Account struct {
	Account      string `json:"account"` // this must be unique per Account record
	Email        string `json:"email"`   // this must be unique per Account record
	PasswordHash []byte `json:"-"`
	//TODO: indicate whether to use password hash or OAuth2, and if OAuth2, which
	//authorization server
}

//FindAccountByID - read a Account record from mongodb by its ID
func FindAccountByID(c *mgo.Collection, id bson.ObjectId) (*Account, error) {
	result := Account{}
	err := c.Find(bson.M{"_id": id}).One(&result)
	return &result, err
}

//FindAccountByEmail - read a Account record from mongodb by the email address
func FindAccountByEmail(c *mgo.Collection, email string) (*Account, error) {
	result := Account{}
	err := c.Find(bson.M{"Email": email}).One(&result)
	return &result, err
}
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Account: %v\n", vars["account"])
	//TODO: error handling, find account, get PUT params, update account
}
