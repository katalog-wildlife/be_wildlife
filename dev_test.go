package be_wildlife

import (
	"fmt"
	"testing"

	model "github.com/katalog-wildlife/be_wildlife/model"
	module "github.com/katalog-wildlife/be_wildlife/module"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	user.Fullname = "Admin"
	user.Email = "admin@gmail.com"
	user.Password = "admin123"
	user.PhoneNumber = "31162877"
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
	user.Email = "admin@gmail.com"
	user.Password = "admin123"
	user, err := module.LogIn(conn, collectionnameUser, user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil LogIn : ", user.Fullname)
	}
}

func TestToken(*testing.T) {
	token := "v4.public.eyJlbWFpbCI6ImFkbWluQGdtYWlsLmNvbSIsImV4cCI6IjIwMjQtMDEtMDlUMTQ6Mzk6MjhaIiwiaWF0IjoiMjAyNC0wMS0wOVQxMjozOToyOFoiLCJpZCI6IjY1OWQzYWRjOTVmOTQwOGQ1ZjJhOWZjNSIsIm5iZiI6IjIwMjQtMDEtMDlUMTI6Mzk6MjhaIn0Q-sAF1CnvrThJBwXD68qSxXCNeeAUhS2FrG0dvXifMnh6f98wKoMs3QVfKtTdN-lnzBcKyKaAmLNIYhNYIrIE"
	tokenstring, err := module.Decode("4300ff0dbdf378880ad20cc7751fb1657583d2c1db9305f6a019dfe30086f00a", token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Id Token : " + tokenstring.Id.Hex())
		fmt.Print("Email Token : " + tokenstring.Email)
	}
}

func TestDeleteAnimal(t *testing.T) {
	col := db
	id := "659691c742b37da5524f2ef6"
	objectId, err := primitive.ObjectIDFromHex(id)
	err = module.DeleteAnimal(objectId, "animal", col)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Hapus Hewan: ")
	}
}

func TestTambahAnimal(t *testing.T) {
	id, err := module.InsertOneDoc(db, "animal", model.Animal{
		ID:             primitive.NewObjectID(),
		Name:           "Harimau Sumatera",
		Namalatin:      "Panthera tigris sumatrae",
		Species:        "Mamalia",
		Habitat:        "Hutan",
		JumlahPopulasi: "400",
		LokasiPopulasi: "Sumatera",
		Status:         "Terancam Punah",
		Description:    "Harimau Sumatera (Panthera tigris sumatrae) adalah subspesies harimau yang hidup di pulau Sumatera dan diakui sebagai subspesies yang terancam punah.",
		Image:          "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/Sumatran_Tiger_2_-_Buffalo_Zoo.jpg/1200px-Sumatran_Tiger_2_-_Buffalo_Zoo.jpg",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Tambah Hewan : ", id)
	}
}

func TestUpdateAnimal(t *testing.T) {
	id := "659b8e396ac7e67935164923"
	objectId, err := primitive.ObjectIDFromHex(id)

	data := module.UpdateOneDoc(objectId, db, "animal", model.Animal{
		ID:             objectId,
		Name:           "Harimau Sumatera",
		Namalatin:      "Panthera tigris sumatrae",
		Species:        "Mamalia",
		Habitat:        "Hutan",
		JumlahPopulasi: "300",
		LokasiPopulasi: "Sumatera",
		Status:         "Terancam Punah",
		Description:    "Harimau Sumatera (Panthera tigris sumatrae) adalah subspesies harimau yang hidup di pulau Sumatera dan diakui sebagai subspesies yang terancam punah.",
		Image:          "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/Sumatran_Tiger_2_-_Buffalo_Zoo.jpg/1200px-Sumatran_Tiger_2_-_Buffalo_Zoo.jpg",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Update Animal", data)
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
