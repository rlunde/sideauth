package main

import (
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*User -- a placeholder for when we build in auth
So far, at least, we're thinking of just having collections tied to
a user, and only having tasks as part of collections. A user might move
a task from one collection to another, or copy a task from one to another,
but it would always be the same user.*/
type User struct {
	// we will just use email as the account identifier for now
	Email        string `json:"email"` // this must be unique per user record
	PasswordHash []byte `json:"-"`
	// Password     string `json:"password,omitempty"`
	//TODO: indicate whether to use password hash or OAuth2, and if OAuth2, which
	//authorization server
}

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

//Crypt use bcrypt to create the password hash
func Crypt(password []byte) ([]byte, error) {
	defer clear(password)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

//FindUserByID - read a user record from mongodb by its ID
func FindUserByID(c *mgo.Collection, id bson.ObjectId) (*User, error) {
	result := User{}
	err := c.Find(bson.M{"_id": id}).One(&result)
	return &result, err
}

//FindUserByEmail - read a user record from mongodb by the email address
func FindUserByEmail(c *mgo.Collection, email string) (*User, error) {
	result := User{}
	err := c.Find(bson.M{"Email": email}).One(&result)
	return &result, err
}

//CreateObjectIDStr - return an ID for use in MongoDB for a user
func CreateObjectIDStr() string {
	var id bson.ObjectId
	id = bson.NewObjectId()
	idstr := id.Hex()
	return idstr
}

//ObjectIDFromIDStr - convert a user ID into a string that can be used in other places
func ObjectIDFromIDStr(idStr string) bson.ObjectId {
	id := bson.ObjectIdHex(idStr)
	return id
}
