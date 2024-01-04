package be_wildlife

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Fullname    string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	PhoneNumber string             `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Salt        string             `json:"salt,omitempty" bson:"salt,omitempty"`
}

type UpdatePassword struct {
	Oldpassword string `json:"oldpassword,omitempty" bson:"oldpassword,omitempty"`
	Newpassword string `json:"newpassword,omitempty" bson:"newpassword,omitempty"`
}

type Animal struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Namalatin      string             `json:"namalatin,omitempty" bson:"namalatin,omitempty"`
	Species        string             `json:"species,omitempty" bson:"species,omitempty"`
	Habitat        string             `json:"habitat,omitempty" bson:"habitat,omitempty"`
	JumlahPopulasi string             `json:"jumlahpopulasi,omitempty" bson:"jumlahpopulasi,omitempty"`
	LokasiPopulasi string             `json:"lokasipopulasi,omitempty" bson:"lokasipopulasi,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
	Image          string             `json:"image,omitempty" bson:"image,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
}

type Credential struct {
	Status  int    `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Response struct {
	Status  int    `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Payload struct {
	Id    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
	Exp   time.Time          `json:"exp"`
	Iat   time.Time          `json:"iat"`
	Nbf   time.Time          `json:"nbf"`
}
