package middlewares

import (
	"net/http"

	"github.com/agomezguru/cti-b2b/auth-service/routers"
)

/* VerifyJWT check if request JWT is valid.
 * This middleware splits the APP between public and private sections
 */
func VerifyJWT(next http.HandlerFunc) http.HandlerFunc  {
	return func (w http.ResponseWriter, r *http.Request)  {
		algo := r.Header.Get("Authorization")
		_, _, _, err := routers.ProcessToken(algo)
		if err != nil {
			http.Error(w, "Token Error: " + err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}