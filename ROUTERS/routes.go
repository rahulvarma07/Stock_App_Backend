package routers

import (
	"encoding/json"
	"net/http"
	controllers "rahulvarma07/github.com/CONTROLLERS"
	models "rahulvarma07/github.com/MODELS"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func LoginTheUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close() // Closing at the end

	var userEnteredAuthModel models.LoginModel

	// Getting all the user entered..
	json.NewDecoder(r.Body).Decode(&userEnteredAuthModel)

	response, err := controllers.LoginTheUser(&userEnteredAuthModel)

	// Checking if there's an error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Bad Request
	} else {
		w.WriteHeader(200) // Success
	}
	defer json.NewEncoder(w).Encode(response)
}

func SignUpTheUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	defer r.Body.Close() // Closing the body

	var userEnteredAuthModel *models.LoginModel

	json.NewDecoder(r.Body).Decode(&userEnteredAuthModel) // Decoding the user entered model

	response, err := controllers.SignUpTheUser(userEnteredAuthModel)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(200)
	}

	defer json.NewEncoder(w).Encode(response)
}
