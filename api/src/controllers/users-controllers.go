package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repo := repositories.NewUserRepository(db)
	repo.Create(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuários!"))
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuários por id!"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
