package controllers

import (
	"net/http"
	"encoding/json"
	u "backend/utils"
	"backend/models"
	"strconv"
	"github.com/gorilla/mux"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(422, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(422, "There was an error in your request"))
		return
	}

	data := models.GetContacts(uint(id))
	resp := u.Message(200, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
