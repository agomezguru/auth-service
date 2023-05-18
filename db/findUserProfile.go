package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/agomezguru/cti-b2b/auth-service/models"
)

func FindUserProfile(ID string) (models.User, error) {

	// Avoids wait for more than specified time
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var userProfile models.User

	db := DatabaseCN

	// Always avoid return user password
	users, err := db.QueryContext(ctx, "SELECT id, login, name, surname, email, photo, status, created_at FROM users WHERE id='"+ID+"'")

	if err != nil {
		fmt.Println("Error reading database: " + err.Error())
		// Here userProfile is empty
		return userProfile, err
	}

	// User finded
	if users.Next() {
		var id int64
		var login, name, surname, email, photo string
		var status int8
		var created time.Time

		err = users.Scan(
			&id,
			&login,
			&name,
			&surname,
			&email,
			&photo,
			&status,
			&created,
		)
		if err != nil {
			// Here userProfile is empty
			return userProfile, err
		}

		userProfile.ID = id
		userProfile.Login = login
		userProfile.Name = name
		userProfile.Surname = surname
		userProfile.Email = email
		userProfile.Photo = photo
		userProfile.Status = status
		userProfile.Created = created

		return userProfile, nil
	}

	return userProfile, errors.New("user register not foundeed")
}
