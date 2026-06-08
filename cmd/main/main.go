package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anaard/simple-student-management/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterStudentManagementRoutes(r)

	http.Handle("/", r)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":" + port, r))
}
