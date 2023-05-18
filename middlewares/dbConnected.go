package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/agomezguru/cti-b2b/auth-service/db"
)

/* This middleware test the DB to check if connections persist */
func DBConnected(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.DBConnectionAlive() == false {
			w.WriteHeader(http.StatusInternalServerError )
			errDescription := "DB connection lost."
			errMap := map[string]string {"message": errDescription}
			message := map[string]interface{}{"error": errMap}
			json.NewEncoder(w).Encode(message)
			return
		}
		// Middlewares always return the same type of input received
		next.ServeHTTP(w, r)
	}
}
