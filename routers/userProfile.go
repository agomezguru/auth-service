package routers

import (
	"encoding/json"
	"net/http"

	"github.com/agomezguru/cti-b2b/auth-service/db"
)

func UserProfile(w http.ResponseWriter, r * http.Request) {
	ID := r.URL.Query().Get("id")

	// Check if id was sended.
	if len(ID) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		errDescription := "User id parameter should be send."
		errMap := map[string]string {"message": errDescription}
		message := map[string]interface{}{"error": errMap}
		json.NewEncoder(w).Encode(message)
		return
	}

	// Locate user id in DB.
	userProfile, err := db.FindUserProfile(ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errDescription := "Error reading database: " + err.Error()
		errMap := map[string]string {"message": errDescription}
		message := map[string]interface{}{"error": errMap}
		json.NewEncoder(w).Encode(message)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userProfile)
	return
}