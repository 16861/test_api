package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//App struct that holds mux Router
type App struct {
	Router *mux.Router
}

//Person stores id, first, lastname and address of person
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address stores cite and state
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

func (a *App) Init() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Anjela", Lastname: "Doe", Address: &Address{City: "City Y", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/people", GetPeople).Methods("GET")
	a.Router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	a.Router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	a.Router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
}

//Run the application on port 8090
func (a *App) Run() {
	if a.Router == nil {
		a.Init()
	}
	http.ListenAndServe(":9090", a.Router)
}

var people []Person

//GetPeople returns all people records
func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(people)
}

//GetPerson returns single person data by id
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//CreatePerson creates person by recieved in post request data
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(people)
}

//DeletePerson deletes a person by id from people
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(people)
	}
}
