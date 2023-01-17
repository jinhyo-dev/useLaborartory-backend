package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Admin struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminSignUp struct {
	Name     string `json:"email"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (u *Admin) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	log.Println(plainText)
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
