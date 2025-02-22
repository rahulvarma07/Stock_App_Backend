package models

type LoginModel struct {
	Email    string `json:"email" validate:"required,emial" bson:"email"`
	Password string `json:"password" bson:"email"`
}
