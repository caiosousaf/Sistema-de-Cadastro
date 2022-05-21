package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Contato struct {
	ID       int 	`json:id`
	Nome     string `json:nome`
	Telefone string `json:telefone`
	Email    string `json:email`
}

var contatos []Contato

func main() {

	rota := mux.NewRouter()

	contatos = append(contatos, Contato{ID: 1, Nome: "João", Telefone: "(88) 2222-2222", Email: "teste123@email.com" })

	rota.HandleFunc("/contatos", getContatos).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", rota))
}

func getContatos(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(contatos)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}