package be_wildlife

import (
	"fmt"
	"testing"

	model "github.com/katalog-wildlife/be_wildlife/model"
	module "github.com/katalog-wildlife/be_wildlife/module"
)

var db = module.MongoConnect("MONGOSTRING", "db_wildlife")
var collectionnameUser = "user"
var collectionnameAnimal = "animal"

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("privateKey : ", privateKey)
	fmt.Println("publicKey : ", publicKey)
}

func TestSignUp(t *testing.T) {
	conn := db
	var user model.User
	user.Fullname = "dito"
	user.Email = "dito@gmail.com"
	user.Password = "dito12345678"
	user.PhoneNumber = "6285718177810"
	email, err := module.SignUp(conn, collectionnameUser, user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil SignUp : ", email)
	}
}

func TestLogIn(t *testing.T) {
	conn := db
	var user model.User
	user.Email = "dito@gmail.com"
	user.Password = "dito12345678"
	user, err := module.LogIn(conn, collectionnameUser, user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil LogIn : ", user.Fullname)
	}
}

// func TestTambahAnimal(t *testing.T) {
// 	conn := db
// 	var wildlife model.Animal
// 	wildlife.Name = "Harimau Sumatera"
// 	wildlife.Namalatin = "Panthera tigris sumatrae"
// 	wildlife.Species = "Mamalia"
// 	wildlife.Habitat = "Hutan"
// 	wildlife.JumlahPopulasi = "400"
// 	wildlife.LokasiPopulasi = "Sumatera"
// 	wildlife.Status = "Terancam Punah"
// 	wildlife.Image = "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/Sumatran_Tiger_2_-_Buffalo_Zoo.jpg/1200px-Sumatran_Tiger_2_-_Buffalo_Zoo.jpg"

// 	//  Perbaikan #1: Memastikan tipe data return dan argumen yang benar
//     _, err := module.PostAnimal(conn, collectionnameAnimal, &wildlife)
//     if err != nil {
//         fmt.Println(err)
//     } else {
//         fmt.Println("Berhasil TambahAnimal : ")
//     }

// }
