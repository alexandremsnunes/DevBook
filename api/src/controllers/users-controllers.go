package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
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
