package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gorilla/mux"
//	"github.com/soura49/restapi/app/src/amster"
	"amster"
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
		fmt.Fprintf(w, "unable to insert the data")
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	out, err := amster.SelectOperationAll()
	if err != nil {
		fmt.Fprintf(w, "Select operation is not working")
	}
	for out.Next() {
		var uuid string
		var info string
		err = out.Scan(&uuid, &info)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(info))
	}
}

func getPeopleByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var info string
	out, err := amster.SelectOperationByID(params["id"])
	if err != nil {
		fmt.Fprintf(w, "It is already Deleted")
	}
	err = out.Scan(&info)
	if err != nil {
		fmt.Fprintf(w, "It is already Deleted")
	}
	w.Write([]byte(info))
}

func deletePeopleByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	out, err := amster.DeleteOperationByID(params["id"])
	if err != nil {
		fmt.Fprintf(w, "It is already Deleted")
	}
	ou := strconv.FormatInt(out, 10)
	w.Write([]byte("Number of Rows Deleted: " + ou))
}

func putPeopleByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ppbi amster.People
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &ppbi)
	out, err := amster.UpdateOperationByID(params["id"], ppbi)
	if err != nil {
		fmt.Fprintf(w, "Issue with the updating operation")
	}
	ou := strconv.FormatInt(out, 10)
	w.Write([]byte("Number of Rows updated: " + ou))
	json.NewEncoder(w).Encode(p)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/people", updatePeople).Methods("POST")
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPeopleByID).Methods("GET")
	router.HandleFunc("/people/{id}", deletePeopleByID).Methods("DELETE")
	router.HandleFunc("/people/{id}", putPeopleByID).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8088", router))
}
