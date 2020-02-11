package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/soura49/restapi/amster"
)

var p amster.People

var pp []amster.People

func generateUUID() (string, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return " ", err
	}
	ou := string(out)
	ou = ou[:len(out)-1]
	return ou, nil
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func updatePeople(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	p.UUID, err = generateUUID()
	if err != nil {
		fmt.Fprintf(w, "unable to create UUID")
	}
	json.Unmarshal(reqBody, &p)
	pp = append(pp, p)
	_, err = amster.InsertOperation(p.UUID, p)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pp)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/people", updatePeople).Methods("POST")
	router.HandleFunc("/people", getPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8088", router))
}