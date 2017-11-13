package main

import (
	"gopkg.in/mgo.v2/bson"
)

//CreateObjectIDStr - return an ID for use in MongoDB for a Account
func CreateObjectIDStr() string {
	var id bson.ObjectId
	id = bson.NewObjectId()
	idstr := id.Hex()
	return idstr
}

//ObjectIDFromIDStr - convert a string into an ID that can be used in other places
func ObjectIDFromIDStr(idStr string) bson.ObjectId {
	id := bson.ObjectIdHex(idStr)
	return id
}
