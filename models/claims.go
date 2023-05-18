package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

/* Claims model Db for JSON web token (JWT) claim
 * Refs.: https://jwt.io
 * 	https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-token-claims
 */
type Claims struct {
	Login string 					`json:"login"`
	ID int64							`json:"id"`
	// Includes in us struct all standard information JWT
	jwt.StandardClaims
}
