package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// Function to hash the password which can be used to store in backend
func HashThePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// classic error checking
	if err != nil {
		return "There is an error in hashing the password", err
	}
	// Successfully hasedThePassword
	return string(hashedPassword), nil
}

// Function to compare password
func CompareThePassword(hashedPassWord, password string) bool {
	checkForError := bcrypt.CompareHashAndPassword([]byte(hashedPassWord), []byte(password))
	// checkForError returns errors if any
	return checkForError == nil
}
