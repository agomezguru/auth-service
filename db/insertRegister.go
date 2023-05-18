package db

import (
	"fmt"
	"time"

	"github.com/agomezguru/cti-b2b/auth-service/models"
)

/* Insert register inside DB
 * Return ID of register inserted in json (int64)
 * Insert successful (bool)
 * Error code
 */
func InsertRegister(u models.User) (bool, error) {
	
	db := DatabaseCN
	
	// Encrypt password before save the new user
	u.Password, _ = EncryptData(u.Password)

	insertNewUser, err := db.Prepare("INSERT INTO users(login, name, surname, email, photo, status, password, created_at, updated_at) VALUES(?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return false, err
	}

	//actualDate := time.Now().Format(time.RFC850)
	_, err = time.LoadLocation("America/Mexico_City")
	if err != nil { // Always check errors even if they should not happen.
		return false, err
	}

	fmt.Print("Intentando insertar registro")
	_, err = insertNewUser.Exec(u.Login, u.Name, u.Surname, u.Email, u.Photo, 1, u.Password, time.Now(), time.Now())

	if err != nil {
		fmt.Print("Hubo un error")
		return false, err
	}
	return true, err
}
