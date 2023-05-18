package utils

import (
	"time"

	"github.com/agomezguru/cti-b2b/auth-service/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* CreateJWT makes a new encrypted token with all user info needed to auth.
 * Ref.: https://www.techtarget.com/searchsecurity/definition/passphrase
 */
func CreateJWT(user models.User) (string, error) {
	myKey := []byte("Kaliman&Solín_personajesInfantilesDeRadio/Novelas")
	
	// Remember never use models.user because this have inside user password 
	payload := jwt.MapClaims(map[string]interface{}{
		"id": 			user.ID,
		"login": 		user.Login,
		"name": 		user.Name,
		"surname":	user.Surname,
		"email": 		user.Email,
		"photo": 		user.Photo,
		"status": 	user.Status,
		"updated": 	user.Updated,
		"exp": 			time.Now().Add(time.Minute * 60).Unix(),
	})

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err  // Remeber tokenStr is empty because generated error.
	}

	return tokenStr, nil
}