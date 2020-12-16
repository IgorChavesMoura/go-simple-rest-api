package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(people)

}

func GetPerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for _, item := range people {

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}

	}

	json.NewEncoder(w).Encode(&Person{})

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var person Person

	_ = json.NewDecoder(r.Body).Decode(&person)

	person.ID = params["id"]

	people = append(people, person)

	json.NewEncoder(w).Encode(people)

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for index, item := range people {

		if item.ID == params["id"] {

			people = append(people[:index], people[index+1:]...)

		}

	}

	json.NewEncoder(w).Encode(people)

}
