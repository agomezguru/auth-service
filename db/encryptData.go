package db

import "golang.org/x/crypto/bcrypt"

/* Encrypt any data string received */
func EncryptData (data string) (string, error) {
	// This si 2 exp cost(8) = 256 the times that data is encrypted recursively 
	cost := 8 
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), cost)

	return string(bytes), err
}
