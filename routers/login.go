package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/agomezguru/cti-b2b/auth-service/db"
	"github.com/agomezguru/cti-b2b/auth-service/models"
	"github.com/agomezguru/cti-b2b/auth-service/utils"
)

/* User login routine */
func Login (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	var loginUser models.User 

	err := json.NewDecoder(r.Body).Decode(&loginUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errDescription := "User and/or password failed. " + err.Error()
		errMap := map[string]string {"message": errDescription}
		message := map[string]interface{}{"error": errMap}
		json.NewEncoder(w).Encode(message)
		return
	}

	// Start validations
	
	objUser, userFinded := db.UserLogin(loginUser.Login, loginUser.Password)
	
	if userFinded == false {
		w.WriteHeader(http.StatusBadRequest)
		errDescription := "This credentials do not match our records."
		errMap := map[string]string {"message": errDescription}
		message := map[string]interface{}{"error": errMap}
		json.NewEncoder(w).Encode(message)
		return
	}
	
	// Login valid. Create a new token for this user device.
	jwtKey, err := utils.CreateJWT(objUser)
	if err != nil {
		// Something was wrong with token
		w.WriteHeader(http.StatusBadRequest)
		errDescription := "Token generation failed." + err.Error()
		errMap := map[string]string {"message": errDescription}
		message := map[string]interface{}{"error": errMap}
		json.NewEncoder(w).Encode(message)
		return
	}

	response := models.LoginToken {
		Token : jwtKey,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	// Set inside cookie the expiration time of the token
	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})

}