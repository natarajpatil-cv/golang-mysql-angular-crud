package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func addApproutes(route *mux.Router) {

	setStaticFolder(route)

	route.HandleFunc("/", renderHome)

	route.HandleFunc("/employees", getEmployees).Methods("GET")

	route.HandleFunc("/employees/{id}", getEmployee).Methods("GET")

	route.HandleFunc("/employees", insertEmployee).Methods("POST")

	route.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")

	route.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")

	fmt.Println("Routes are Loded.")
}
