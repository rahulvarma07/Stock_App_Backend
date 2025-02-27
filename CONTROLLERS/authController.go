package controllers

import (
	"context"
	"errors"
	"fmt"

	database "rahulvarma07/github.com/DATABASE"
	helpers "rahulvarma07/github.com/HELPERS"
	models "rahulvarma07/github.com/MODELS"

	"github.com/go-playground/validator"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var MongoClient = database.GetMongoCLient()
var UserMongoCollection *mongo.Collection = database.CreateMongoCollection(MongoClient, "UserData")

// Login func checking for the user...
func LoginTheUser(ref *models.LoginModel) (models.Response, error) {

	var finalResponse models.Response // Response type to be returned

	// check the user
	// 1.Get the mail & Get the password

	// if there is a user
	validate := bson.M{"email": ref.Email} // Trying to find user entered mail

	// type of user model to retrive the data
	type User struct {
		Email    string `bson:"email"`
		Password string `bson:"password"`
	}

	var ExsistingCredential User

	err := UserMongoCollection.FindOne(context.Background(), validate).Decode(&ExsistingCredential)

	fmt.Println(ExsistingCredential)

	if err == mongo.ErrNoDocuments { // Corrected error check
		finalResponse.Status = "UserSideBadStatus"
		finalResponse.Message = "User does not exist in database"
		return finalResponse, err
	}

	// Debugging
	fmt.Printf("Fetched User: %+v\n", ExsistingCredential)
	fmt.Println("Retrieved Password:", ExsistingCredential.Password)

	// If password is still empty, investigate MongoDB data

	isPasswordMatching := helpers.CompareThePassword(ExsistingCredential.Password, ref.Password)
	if !isPasswordMatching {
		finalResponse.Status = "UserSideBadStatus"
		finalResponse.Message = "Enter Valid Password!"
		return finalResponse, errors.New("invalid password")
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

	validate := validator.New()
	checkValidation := validate.Struct(ref)

	if checkValidation != nil {
		finalResponse.Status = "UserSideBadResponse"
		finalResponse.Message = "Invalid mail"
		return finalResponse, errors.New("invalid cred")
	}
	// first check whether the mail is present

	findUserWithEmail := bson.M{"email": ref.Email}
	isUserExsist := UserMongoCollection.FindOne(context.Background(), findUserWithEmail)

	// if it is
	if isUserExsist.Err() != mongo.ErrNoDocuments {
		finalResponse.Status = "UserSideBadResponse"
		finalResponse.Message = "User Email Already Exsists"
		return finalResponse, errors.New("user already exsists")
	}

	// if it's not
	hashUserEnteredPassword, err := helpers.HashThePassword(ref.Password)
	if err != nil {
		finalResponse.Status = "ServerSideBadResponse"
		finalResponse.Message = "Unable to hash the password"
		return finalResponse, err
	}

	ref.Password = hashUserEnteredPassword

	addTheUser, err := UserMongoCollection.InsertOne(context.Background(), ref)
	print(addTheUser)

	if err != nil {
		finalResponse.Status = "ServerSideBadResponse"
		finalResponse.Message = "Unable to add the user"
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
