package routers

import (
	"errors"
	"strings"

	"github.com/agomezguru/cti-b2b/auth-service/db"
	"github.com/agomezguru/cti-b2b/auth-service/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* UserLogin is the user login to use in all endpoints. */
var UserLogin string

/* UserID ID returned by model to use in all endpointsl */
var UserID int64
/* ProcessToken extracts all values inside token sended and return
 * return this values in a JWT Claims model arrangement.
 */
func ProcessToken(tkn string) (*models.Claims, bool, int64, error) {

	myKey := []byte("Kaliman&Solín_personajesInfantilesDeRadio/Novelas")
	
	claims := &models.Claims{}

	splitToken := strings.Split(tkn, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, -1, errors.New("token format is invalid") 
	}

	// Clean's token
	tkn = strings.TrimSpace(splitToken[1])

	token, err := jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token)(interface{}, error){
		return myKey, nil
	})

	if err == nil {
		_, finded, ID := db.UserExist(claims.Login, string(""))

		if finded == true {
			UserLogin = claims.Login
			UserID = claims.ID
		}

		return claims, finded, ID, nil
	}

	if !token.Valid {
		return claims, false, -1, errors.New("invalid token")
	}

	return claims, false, -1, err
}
