package db

import (
	"context"
	"time"

	"github.com/agomezguru/cti-b2b/auth-service/models"
)

/* Verify if user was created before
 * Return object User if exist (models.User)
 * true if exist (bool)
 * User ID if exist (int64) return <> -1
 */
func UserExist(login, email string) (models.User, bool, int64)  {
	
	// Avoids wait for more than specified time
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	var user models.User

	db := DatabaseCN
	
	// First search by login, this must be unique.
	users, err := db.QueryContext(ctx, "SELECT id, login, email, password FROM users WHERE login='" + login + 
		"'  OR email='" + email + "'")

	if err != nil {
		// Here user is empty
		return user, false, -1
	}

	// User finded 
	if users.Next() {
		var id int64
		var login, email, password string
		err = users.Scan(&id, &login, &email, &password)
		if err != nil {
			// Here user is empty
			return user, false, -1
		}
		user.ID = id
		user.Login = login
		user.Email = email
		user.Password = password
		return user, true, user.ID
	}

	// Here user is empty
	return user, false, -1
}
