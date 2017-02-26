package models

import "time"

// Credential is a Model to storage
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User is a Model
type User struct {
	ID         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	CreatedOn  time.Time  `json:"created_on"`
	Credential Credential `json:"credential"`
}
