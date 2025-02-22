package controllers

import (
	"context"
	"errors"
	"log"
	database "rahulvarma07/github.com/DATABASE"
	helpers "rahulvarma07/github.com/HELPERS"
	models "rahulvarma07/github.com/MODELS"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var UserMongoCollection *mongo.Collection = database.CreateMongoCollection(database.GetMongoCLient(), "UserData")

// Login func checking for the user...
func LoginTheUser(ref *models.LoginModel) (models.Response, error) {

	var finalResponse models.Response // Response type to be returned

	// check the user
	// 1.Get the mail & Get the password
	usreEnteredMail := ref.Email
	userEnteredPassword := ref.Password

	// if there is a user
	validate := bson.M{"emial": usreEnteredMail} // Trying to find user entered mail
	var ExsistingCredential struct {
		Password string `json:"password"`
	}
	err := UserMongoCollection.FindOne(context.Background(), validate).Decode(&ExsistingCredential)

	// If There's no user..
	if err != nil {
		finalResponse.Status = "UserSideBadStatus"
		finalResponse.Message = "User does not exsits in database"
		return finalResponse, err
	}

	// If there's a user
	isPasswordMatching := helpers.CompareThePassword(ExsistingCredential.Password, userEnteredPassword)
	if !isPasswordMatching {
		finalResponse.Status = "UserSideBadStatus"
		finalResponse.Message = "Enter Valid Password!"
		return finalResponse, errors.New("inValid Password")
	}

	token, err := helpers.GenerateToken(ref)
	if err != nil {
		finalResponse.Status = "ServerSideBadResponse"
		finalResponse.Message = "Unable to generate a token"
		return finalResponse, errors.New("unable to generate a token")
	} else {
		finalResponse.Status = "SuccesState"
		finalResponse.Message = "Logined The User Successfully"
		finalResponse.TokenString = token
	}
	return finalResponse, nil
}

func SignUpTheUser(ref *models.LoginModel) (models.Response, error) {
	var finalResponse models.Response

	// first check whether the mail is present
	userEnteredEmail := ref.Email
	userEnteredPassword := ref.Password

	findUserWithEmail := bson.M{"email": userEnteredEmail}
	isUserExsist := UserMongoCollection.FindOne(context.Background(), findUserWithEmail)

	// if it is
	if isUserExsist == nil {
		finalResponse.Status = "UserSideBadResponse"
		finalResponse.Message = "User Email Already Exsists"
		return finalResponse, errors.New("user already exsists")
	}

	// if it's not
	hashUserEnteredPassword, err := helpers.HashThePassword(userEnteredPassword)
	if err != nil {
		finalResponse.Status = "ServerSideBadResponse"
		finalResponse.Message = "Unable to hash the password"
		return finalResponse, err
	}

	ref.Password = hashUserEnteredPassword

	addTheUser, err := UserMongoCollection.InsertOne(context.Background(), ref)
	if err != nil {
		finalResponse.Status = "ServerSideBadResponse"
		finalResponse.Message = "Unable to add the user"
		log.Println("unable to add the user with", addTheUser.InsertedID)
		return finalResponse, err
	}

	token, tokenErr := helpers.GenerateToken(ref)

	if tokenErr != nil {
		finalResponse.Status = "ServerSideBadResponse"
		finalResponse.Message = "Unable to genereate the token"
		return finalResponse, err
	} else {
		finalResponse.Status = "Success"
		finalResponse.Message = "Successfully Signed in the user"
		finalResponse.TokenString = token
		return finalResponse, nil
	}
}
