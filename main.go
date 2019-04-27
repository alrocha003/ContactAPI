package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

/*
type Calc struct {
	Number   int
	Percente int
}

type Pessoa struct {
	ID   int
	Nome string
}
*/

type Contact struct {
	Id       int
	Nome     string
	Telefone string
	Email    string
}

//var pessoas []Pessoa
var contacts []Contact

func main() {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/contact", addContact).Methods("POST")
	rotas.HandleFunc("/contact", getAllContacts).Methods("GET")
	rotas.HandleFunc("/contactRemove", removeContact).Methods("POST")
	rotas.HandleFunc("/contactUpdate", updateContact).Methods("POST")
	rotas.HandleFunc("/Random", getRandomNum).Methods("GET")

	var port = ""

	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World da Api :D")
}

func getAllContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
}

func addContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(http.StatusOK)
}

func removeContact(w http.ResponseWriter, r *http.Request) {
	var id int
	_ = json.NewDecoder(r.Body).Decode(&id)
	var newArray []Contact

	for i := 0; i < len(contacts); i++ {
		if id == contacts[i].Id {
			continue
		} else {
			newArray = append(newArray, contacts[i])
		}
	}

	contacts = newArray
	json.NewEncoder(w).Encode(http.StatusOK)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)

	for i := 0; i < len(contacts); i++ {
		if contact.Id == contacts[i].Id {
			contacts[i].Nome = contact.Nome
			contacts[i].Telefone = contact.Telefone
			contacts[i].Email = contact.Email
		} else {
			continue
		}
	}

	json.NewEncoder(w).Encode(http.StatusOK)
}

func getAllH(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Essa é a rota getallh1, não é a principal")
}

func getRandomNum(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rand.Intn(1000))
}

/*
func getPercent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var calc Calc
	_ = json.NewDecoder(r.Body).Decode(&calc)

	var res = (calc.Number * calc.Percente) / 100
	json.NewEncoder(w).Encode(res)
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var p Pessoa
	_ = json.NewDecoder(r.Body).Decode(&p)

	pessoas = append(pessoas, p)
	json.NewEncoder(w).Encode(http.StatusOK)
}

func getAllPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

func removePeople(w http.ResponseWriter, r *http.Request) {
	var id int
	_ = json.NewDecoder(r.Body).Decode(&id)
	var newArray []Pessoa

	for _, v := range a {
		if id == pessoas[i] {
			continue
		} else {
			newArray = append(newArray, v)
		}
	}

	pessoas = newArray
}*/
