package main

import (
	"authentication/cmd/data"
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user againse the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentiasl"), http.StatusBadRequest)
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) Register(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"emai"`
		FistName string `json:"first_name"`
		LastName string `json:"last_name"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	user := data.User{
		Email:     requestPayload.Email,
		FirstName: requestPayload.FistName,
		LastName:  requestPayload.LastName,
		Password:  requestPayload.Password,
		Active:    1,
	}
	user.SetPassword(requestPayload.Password)

	insertedUser, err := app.Models.User.Insert(user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "User registered successfully",
		Data:    insertedUser,
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}
