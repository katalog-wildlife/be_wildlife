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
	user.Fullname = "Ade"
	user.Email = "adecand12@gmail.com"
	user.Password = "12345678"
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
	user.Email = "adecand12@gmail.com"
	user.Password = "12345678"
	user, err := module.LogIn(conn, collectionnameUser, user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil LogIn : ", user.Fullname)
	}
}

func TestToken(*testing.T) {
	token := "v4.public.eyJleHAiOiIyMDI0LTAxLTA0VDExOjI1OjU0WiIsImZ1bGxuYW1lIjoiYWRtaW5AZ21haWwuY29tIiwiaWF0IjoiMjAyNC0wMS0wNFQwOToyNTo1NFoiLCJpZCI6IjY1OTY1ZWNkY2MxOGQxNmNkNGNhNGY4YSIsIm5iZiI6IjIwMjQtMDEtMDRUMDk6MjU6NTRaIn22kA21UMcQv-6lNrkBu88rV3XGGgToTBqulQui3HrZcYb_Go-qyCBdzje7Qg3Omj-hI5lXRRFj1afCzeMdyG0B"
	tokenstring, err := module.Decode("dc9a05bd1679ffa792336245874399894f513a9d38a0f108907c2d6713fc4db5", token)
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
