package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonGreeting struct {
	Person   Person `json:"person"`
	Greeting string `json:"greeting"`
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "William", Age: 42}
	personGreeting := PersonGreeting{
		Person:   person,
		Greeting: fmt.Sprintf("Hello %s (%d)", person.Name, person.Age),
	}
	b, err := json.Marshal(personGreeting)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func createPerson(w http.ResponseWriter, req *http.Request) {
	var personGreeting PersonGreeting
	if err := json.NewDecoder(req.Body).Decode(&personGreeting); err != nil {
		_, _ = fmt.Fprintf(w, "Error parsing JSON: "+err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Printf("Person: %v\nRegistered greeting: %s\n",
		personGreeting.Person, personGreeting.Greeting)
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		getPerson(w, req)
	} else if req.Method == http.MethodPost {
		createPerson(w, req)
	} else {
		_, _ = fmt.Fprintf(w, "Method is not supported")
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func main() {
	http.HandleFunc("/person", handler)

	log.Println("\nRegistered endpoints:\n\t* GET /person\n\t* POST /person\nRunning on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
