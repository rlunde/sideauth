package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

//UpdateAccount - TODO
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Account: %v\n", vars["account"])
	//TODO: error handling, find account, get PUT params, update account
}
