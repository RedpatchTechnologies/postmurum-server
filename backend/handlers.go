package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/RedPatchTechnologies/postmurum-server/backend/models"
	//"github.com/adam-hanna/jwt-auth/jwt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/pop"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var cred Credentials
var conf *oauth2.Config

type AuthTokenResponse struct {
	AuthToken string
}

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

// RandToken generates a random @l length token.
func RandToken(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func init() {
	file, err := ioutil.ReadFile("./env/creds.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &cred)

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://localhost/api/oauthcallback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			//"https://www.googleapis.com/auth/cloud-platform",

			// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

// IndexHandler handels /.
func IndexHandler(c *gin.Context) {

	db, dberr := pop.Connect("development")
	if dberr != nil {
		log.Panic(dberr)
	}

	fmt.Printf("db (index handler) %+v\n", db)

	query := models.DB
	users := []models.Organization{}
	err := query.All(&users)
	fmt.Printf("users is %+v\n", users)

	if err != nil {
		fmt.Printf("fetch all orgs error: %v\n", err)

	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
func AuthHandler(c *gin.Context) {
	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	queryState := c.Request.URL.Query().Get("state")

	if retrievedState != queryState {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
		return
	}

	tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	session.Set("AccessToken", tok.AccessToken)
	session.Set("RefreshToken", tok.RefreshToken)
	session.Set("TokenType", tok.TokenType)
	session.Set("Expiry", tok.Expiry.Format(time.RFC3339))

	client := conf.Client(oauth2.NoContext, tok)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer email.Body.Close()
	data, _ := ioutil.ReadAll(email.Body)
	log.Println("Email body: ", string(data))

	token := session.Get("secureToken")
	log.Println("token is: ", token)

	var authTokenForUser = "somesortofauthtoken"

	session.Set(token, authTokenForUser)
	session.Save()
	c.Redirect(http.StatusFound, "http://localhost/finallogin")
}

// LoginHandler handles the login procedure.
func LoginHandler(c *gin.Context) {

	token := c.Request.URL.Query().Get("token")

	state := RandToken(32)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Set("secureToken", token)
	session.Save()
	log.Printf("Login: Stored session: %v\n", state)
	link := getLoginURL(state)
	c.HTML(http.StatusOK, "auth.tmpl", gin.H{"link": link})
}

// FieldHandler is a rudementary handler for logged in users.
func FieldHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "field.tmpl", gin.H{"user": userID})
}

func SubmitHandler(c *gin.Context) {
	state := RandToken(32)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	log.Printf("Login: Stored session: %v\n", state)
	link := getLoginURL(state)
	c.HTML(http.StatusOK, "auth.tmpl", gin.H{"link": link})
}

func OrgHandler(c *gin.Context) {
	db, dberr := pop.Connect("development")
	if dberr != nil {
		log.Panic(dberr)
	}

	fmt.Printf("db (index handler) %+v\n", db)

	query := models.DB
	users := []models.Organization{}
	err := query.All(&users)
	fmt.Printf("users is %+v\n", users)

	c.Header("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Printf("fetch all orgs error: %v\n", err)

	}
	c.JSON(http.StatusOK, users)

	//c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func AuthTokenHandler(c *gin.Context) {

	token := c.Request.URL.Query().Get("token")
	fmt.Printf("AuthTokenHandler token is %+v\n", token)
	session := sessions.Default(c)
	fmt.Printf("AuthTokenHandler session is %+v\n", session)
	authToken := session.Get(token)
	fmt.Printf("AuthTokenHandler authToken is %+v\n", authToken)
	//session.Set("authtoken", authTokenForUser)

	//response := `{"authtoken":"authToken"}`

	authTokenResponse := AuthTokenResponse{}
	authTokenResponse.AuthToken = authToken.(string)
	fmt.Printf("AuthTokenHandler authTokenResponse is %+v\n", authTokenResponse)

	c.JSON(http.StatusOK, authTokenResponse)
}
