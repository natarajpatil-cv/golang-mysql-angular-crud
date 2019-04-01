package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	fmt.Println("Server will start at http://localhost:8000/")

	connectDatabse()

	route := mux.NewRouter()

	addApproutes(route)

	// Use default options
	handler := cors.AllowAll().Handler(route)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
