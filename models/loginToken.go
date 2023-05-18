package models

/* LoginToken contains token returned by login func */
type LoginToken struct {
	Token string `json:"token,omitempty"`
}
