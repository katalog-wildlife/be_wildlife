package be_wildlife

import (
	"encoding/json"
	"net/http"

	model "github.com/katalog-wildlife/be_wildlife/model"
	module "github.com/katalog-wildlife/be_wildlife/module"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	credential model.Credential
	response   model.Response
	user       model.User
	password   model.UpdatePassword
)

func SignUpHandler(MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := module.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return module.GCFReturnStruct(response)
	}
	email, err := module.SignUp(conn, collectionname, user)
	if err != nil {
		response.Message = err.Error()
		return module.GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil SignUp"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data": bson.M{
			"email": email,
		},
	}
	return module.GCFReturnStruct(responData)
}
