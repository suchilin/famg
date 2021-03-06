package controllers

import (
	"backend/app"
	"backend/models"
	u "backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(422, "Invalid request"))
		return
	}

	resp := account.Create()

	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(422, "Invalid request"))
		return
	}

	resp, ok := models.Login(account.Username, account.Password)

	if ok {
		app.SetAccessCookie(account.Username, account.ID, w)
		app.SetRefreshCookie(account.Username, account.ID, w)
	}
	u.Respond(w, resp)
}

var IsAuthenticated = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("is authenticated")
	u.Respond(w, u.Message(http.StatusOK, "is authenticated"))
}

var SignOut = func(w http.ResponseWriter, r *http.Request){
	app.ClearSession(w)
	u.Message(200, "Log Out sucefull")
}
