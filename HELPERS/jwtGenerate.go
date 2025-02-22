package helpers

import (
	"log"
	"os"

	models "rahulvarma07/github.com/MODELS"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	// Classic error checking
	if err != nil {
		log.Fatal("There is an error in loading ENV-{HELPERS/TOKEN}")
	}
}

// Secret Key ~
var SecretKey = []byte(os.Getenv("SECRET_KEY"))

// function to generate key
func GenerateToken(userDetails *models.LoginModel) (string, error) {
	// What contents should be displayed while decoding the token..
	claims := jwt.MapClaims{
		"email": userDetails.Email,
	}
	// Using HS256 for generating the key..
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)

	// Classic error checking
	if err != nil {
		return "Error in creating a JWT token", err
	}
	return tokenString, nil
}
