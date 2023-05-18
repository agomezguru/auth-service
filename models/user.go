package models

import (
	"time"
)

/* User model for MySQL DB */
type User struct {
	ID int64							`json:"id"`
	Login string 					`json:"login"`
	Name string 					`json:"name"`
	Surname string 				`json:"surname"`
	Email string 					`json:"email"`
	Photo string 					`json:"photo,omitempty"`
	Status int8 					`json:"status,omitempty"`
	Email_ok time.Time 		`json:"verified,omitempty"`
	Password string 			`json:"password"`
	Remember_token string `json:"token,omitempty"`
	Created time.Time 	  `json:"created,omitempty"`
	Updated time.Time 		`json:"updated,omitempty"`
}	

/*
+-------------------+-----------------+------+-----+---------+----------------+
| Field             | Type            | Null | Key | Default | Extra          |
+-------------------+-----------------+------+-----+---------+----------------+
| id                | bigint unsigned | NO   | PRI | NULL    | auto_increment |
| login             | varchar(191)    | NO   | UNI | NULL    |                |
| name              | varchar(191)    | NO   |     | NULL    |                |
| surname           | varchar(191)    | NO   |     | NULL    |                |
| email             | varchar(191)    | NO   | UNI | NULL    |                |
| photo             | varchar(191)    | NO   |     | NULL    |                |
| status            | tinyint(1)      | NO   |     | NULL    |                |
| email_verified_at | timestamp       | YES  |     | NULL    |                |
| password          | varchar(191)    | NO   |     | NULL    |                |
| remember_token    | varchar(100)    | YES  |     | NULL    |                |
| created_at        | timestamp       | YES  |     | NULL    |                |
| updated_at        | timestamp       | YES  |     | NULL    |                |
+-------------------+-----------------+------+-----+---------+----------------+
*/
