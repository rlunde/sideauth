package main

import (
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*Account -- a unique name correcponding to (for now) a password hash, and that can
  have sessions.
*/
type Account struct {
	Account      string `json:"account"` // this must be unique per Account record
	Email        string `json:"email"`   // this must be unique per Account record
	PasswordHash []byte `json:"-"`
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

//CreateObjectIDStr - return an ID for use in MongoDB for a Account
func CreateObjectIDStr() string {
	var id bson.ObjectId
	id = bson.NewObjectId()
	idstr := id.Hex()
	return idstr
}

//ObjectIDFromIDStr - convert a Account ID into a string that can be used in other places
func ObjectIDFromIDStr(idStr string) bson.ObjectId {
	id := bson.ObjectIdHex(idStr)
	return id
}
