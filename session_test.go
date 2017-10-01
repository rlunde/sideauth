package main

import (
	"testing"
)

var (
	DropDatabase = true
)

func SetupTestSession() *Session {
	mgr := GetMgr()
	idstr := mgr.sessionID()
	session := NewSession(mgr, idstr)
	return &session
}
func TestCreateReadDestroy(t *testing.T) {
	mgr := GetMgr()
	err := GetDatabaseConnection(mgr)
	session := SetupTestSession()

	//Test Create, Set, and Get
	err = Create(session)
	if err != nil {
		t.Errorf("Create failed: %s\n", err.Error())
	}
	// add something to the map
	session.Set("email", "al@pa.ca")
	email := session.Get("email")
	if email != "al@pa.ca" {
		t.Errorf("session should have email: %s but has email: %s", "al@pa.ca", email)
	}

	//Test Read and Get
	returnedSession := NewSession(session.Mgr, session.SessionID)
	err = Read(&returnedSession)
	if err != nil {
		t.Errorf("Read failed: %s\n", err.Error())
	}
	if session.SessionID != returnedSession.SessionID {
		t.Errorf("Create has SessionID: %s but returnedSession has SessionID: %s", session.SessionID, returnedSession.SessionID)
	}
	returnedEmail := returnedSession.Get("email")
	if returnedEmail != email {
		t.Errorf("loaded session has email: %s but should have email: %s", returnedEmail, email)
	}

	//Test Set (which does an Update), Read, and Get
	returnedSession.Set("email", "al@cap.one")
	returnedSession.Set("foo", "bar") // for testing Delete
	session3 := NewSession(session.Mgr, session.SessionID)
	err = Read(&session3)
	if err != nil {
		t.Errorf("Read failed: %s\n", err.Error())
	}
	if session.SessionID != session3.SessionID {
		t.Errorf("Create has SessionID: %s but session3 has SessionID: %s", session.SessionID, session3.SessionID)
	}
	email3 := session3.Get("email")
	returnedEmail = returnedSession.Get("email")
	if returnedEmail != email3 {
		t.Errorf("saved session has email: %s but re-read session email: %s", returnedEmail, email3)
	}
	foo := session3.Get("foo")
	if foo != "bar" {
		t.Errorf("expected key foo to have value bar, but got %s", foo)
	}
	session3.Delete("foo")
	session4 := NewSession(session.Mgr, session.SessionID)
	err = Read(&session4)
	foo2 := session4.Get("foo")
	if foo2 != nil {
		t.Errorf("expected key foo to have value nil, but got %s", foo2)
	}

	err = Destroy(session)
	if err != nil {
		t.Errorf("Destroy failed: %s\n", err.Error())
	}
	goneSession := NewSession(session.Mgr, session.SessionID)
	err = Read(&goneSession)
	if err == nil || err.Error() != "not found" {
		t.Errorf("Read found a deleted session: %s\n", session.SessionID)
	}
}
