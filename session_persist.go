package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/* Note: see article on escape analysis: references should be passed into functions, not out.
http://www.agardner.me/golang/garbage/collection/gc/escape/analysis/2015/10/18/go-escape-analysis.html
*/

//Config -- keep track of all the config data
type Config struct {
	MongoSession        *mgo.Session    `json:"-"`
	MongoCollection     *mgo.Collection `json:"-"`
	MongoHost           string          `json:"host"`
	MongoDatabase       string          `json:"database"`
	MongoCollectionName string          `json:"collection"`
}

//LoadConfiguration --load the configuration values from session.cfg
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

//DbConn -- return the database connection
func (mgr *Manager) DbConn() *mgo.Collection {
	return mgr.sessionConfig.MongoCollection
}

//GetSessionConfig -- return the config data for the session
func GetSessionConfig(mgr *Manager) {
	sessionConfig := LoadConfiguration("session.cfg")
	mgr.sessionConfig = sessionConfig
}

//GetDatabaseConnection - open a Mongo database for storing sessions
func GetDatabaseConnection(mgr *Manager) (err error) {
	//Note: pass globalSessionMgr as the argument to this function
	if mgr == nil {
		err = errors.New("GetDatabaseConnection called with nil Manager")
		return
	}
	MongoSession, err := mgo.Dial(mgr.sessionConfig.MongoHost)
	if err != nil {
		fmt.Printf("Could not open mongo database session: %s", err.Error())
		return err
	}
	mgr.sessionConfig.MongoSession = MongoSession
	mgr.sessionConfig.MongoSession.SetMode(mgo.Monotonic, true)
	// Error check on every access
	mgr.sessionConfig.MongoSession.SetSafe(&mgo.Safe{})

	MongoCollection := mgr.sessionConfig.MongoSession.DB(mgr.sessionConfig.MongoDatabase).C(mgr.sessionConfig.MongoCollectionName)
	mgr.sessionConfig.MongoCollection = MongoCollection
	return
}

//Create - create a new session record in MongoDB
func Create(session *Session) (err error) {
	mgr, err := checkMgrAndSession(session, "Create")
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	// session.ID = bson.NewObjectId()
	//c := mgr.sessionConfig.MongoCollection // TODO: figure out why this doesn't work
	c := mgr.sessionConfig.MongoSession.DB("test").C("sessions")
	err = c.Insert(session)
	if err != nil {
		return err
	}
	return nil
}

func checkMgrAndSession(session *Session, fn string) (mgr *Manager, err error) {
	if session == nil {
		err = errors.New(fn + " called with nil Session")
		return nil, err
	}
	mgr = session.Mgr
	if mgr == nil {
		err = errors.New(fn + " called with nil session.Mgr")
		return
	}
	if mgr.sessionConfig.MongoSession == nil {
		err = errors.New(fn + " called with nil Manager MongoSession")
	}
	return
}

//Read -- get the session out of mongodb
func Read(session *Session) (err error) {
	mgr, err := checkMgrAndSession(session, "Read")
	if err != nil {
		return err
	}
	c := mgr.sessionConfig.MongoSession.DB("test").C("sessions")

	err = c.Find(bson.M{"sessionid": session.SessionID}).One(session)
	session.Mgr = mgr // mongo wipes out the struct before creating a new one
	return err        // err is nil if it found it
}

//Destroy -- delete a session record from mongodb
func Destroy(session *Session) (err error) {
	mgr, err := checkMgrAndSession(session, "Destroy")
	if err != nil {
		return err
	}
	c := mgr.sessionConfig.MongoSession.DB("test").C("sessions")

	//c := mgr.sessionConfig.MongoCollection
	err = c.Remove(bson.M{"sessionid": session.SessionID})
	return err // err is nil if it found it
}

//Update -- update a session in mongodb, and update the last access time
func Update(session *Session) (err error) {
	mgr, err := checkMgrAndSession(session, "Update")
	if err != nil {
		return err
	}
	c := mgr.sessionConfig.MongoSession.DB("test").C("sessions")

	//c := mgr.sessionConfig.MongoCollection
	err = c.Update(bson.M{"sessionid": session.SessionID}, session)
	session.Mgr = mgr // mongo may wipe out the struct before creating a new one

	return err // err is nil if it found it
}

func init() {

}
