package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetCollection -- get the accounts collection from the localhost database
func GetCollection() (c *mgo.Collection) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("GetCollection error: %s", err) // TODO: figure out what to do
		return nil
	}
	c = session.DB("sideauth").C("accounts")
	return c
}

//FindAccountByName - read a Account record from mongodb by its account name
func FindAccountByName(c *mgo.Collection, name string) (*Account, error) {
	result := Account{}
	err := c.Find(bson.M{"Name": name}).One(&result)
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
