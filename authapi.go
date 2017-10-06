package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	// TODO: use regular http and gorilla rather than gin
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

/*RunService runs the main service endpoints
 */
func RunService() {
	// next 3 lines show use of Gorialla mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":10000", router))

	// next steps: replace all the Gin stuff with net/http + gorilla
	r := gin.Default()
	r.Static("/js", "../client/js")
	r.Static("/css", "../client/css")
	r.Static("/img", "../client/img")

	r.LoadHTMLGlob("../client/*.html")

	r.POST("/register", RegisterAccount)
	/* session related operations: login creates a session, logout destroys one */
	r.POST("/login", LoginWithAccount)
	r.POST("/logout", Logout)

	/* all other operations require a valid session, and validation happens as a first step */

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// TODO: check r.Run for error return
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

//RegisterAccount -- create a new login
func RegisterAccount(c *gin.Context) {
	username, email, password, err := getRegistrationData(c)
	if err != nil {
		c.AbortWithError(400, err)
	} else {
		fmt.Printf("RegisterAccount called with username %s, email %s, password %s\n", username, email, password)
	}
	//TODO: validate that account doesn't already exist
	//TODO: try to create login and save it in database
	//TODO: create a session cookie
	//TODO: return success or error message
	//TODO: on success, send email and display a verify email form
	//TODO: on error, display error message and redirect to register form
}

//LoginWithAccount -- create a new session, or return an error
func LoginWithAccount(c *gin.Context) {
	username, password, err := getLoginData(c)
	if err != nil {
		c.AbortWithError(400, err)
	} else {
		fmt.Printf("LoginWithAccount called with username %s, password %s\n", username, password)
	}
	//TODO: create a session cookie
	w := c.Writer
	r := c.Request
	mgr := GetMgr()
	sess, err := mgr.SessionStart(w, r)
	if err != nil {
		c.AbortWithError(400, err)
	}
	//TODO: return success or error message
	//TODO: verify password is correct
	sess.Set("username", username)
	http.Redirect(w, r, "/", 302)
	//TODO: on error, display error message and redirect back to login form
}

//Logout -- destroy a session
func Logout(c *gin.Context) {
	w := c.Writer
	r := c.Request
	mgr := GetMgr()
	err := mgr.SessionEnd(w, r)
	if err != nil {
		c.AbortWithError(400, err)
	}
	http.Redirect(w, r, "/", 302)
	//TODO: return success or error message
	//TODO: on error, display error message and redirect back to login form
}

/*Registration - need to use BindJSON to retrieve from gin, since now posting from React as JSON struct */
type Registration struct {
	Username     string `form:"username" json:"username" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	ConfPassword string `form:"confpassword" json:"confpassword" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	Remember     bool   `form:"remember" json:"remember" `
}

func getRegistrationData(c *gin.Context) (username, email, password string, err error) {

	var json Registration
	err = c.BindJSON(&json)
	if err == nil {
		fmt.Printf("Got username: %s\n", json.Username)
	}
	err = checkmail.ValidateFormat(json.Email)
	if err != nil {
		return
	}
	// err = checkmail.ValidateHost(json.Email)
	// if err != nil {
	// 	return
	// }
	if json.Password != json.ConfPassword {
		err = errors.New("Password and confirm-password do not match")
	}
	username = json.Username
	email = json.Email
	password = json.Password
	return
}

/*Login - need to use BindJSON to retrieve from gin, since now posting from React as JSON struct */
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Remember bool   `form:"remember" json:"remember" `
}

func getLoginData(c *gin.Context) (username, password string, err error) {

	var json Login
	err = c.BindJSON(&json)
	if err == nil {
		fmt.Printf("Got username: %s\n", json.Username)
	}
	username = json.Username
	password = json.Password
	return
}
