package db

import (
	"github.com/agomezguru/cti-b2b/auth-service/models"
	"golang.org/x/crypto/bcrypt"
)

/* UserLogin verify if user is registred  in DB and password is valid. */
func UserLogin (login string, password string) (models.User, bool) {
	
	user, finded, _ := UserExist(login, string(""))

	if finded == false {
		// Here user is empty
		return user, false
	}
	
	// Verify if password typed is valid.
	loginPassword := []byte(password)
	dbSavedPassword := []byte(user.Password)
	
	err := bcrypt.CompareHashAndPassword(dbSavedPassword, loginPassword)

	if err != nil {
		// Here user must be empty
		return user, false
	}

	return user, true
}
