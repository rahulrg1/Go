package main

import (
	"fmt"
	"go_valid/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.GetAllData).Methods("GET")
	r.HandleFunc("/add", middleware.CreateForm).Methods("POST")
	r.HandleFunc("/update/{email}", middleware.UpdateForm).Methods("PUT")
	r.HandleFunc("/delete/{email}", middleware.DeleteForm).Methods("DELETE")

	router := r.PathPrefix("/api").Subrouter()
	router.HandleFunc("/{email}", middleware.GetDataByEmail).Methods("GET")
	router.HandleFunc("/", middleware.GetDataByCity).Queries("location", "{location}").Methods("GET")

	fmt.Println("server started at 8000 !")
	log.Fatal(http.ListenAndServe(":8000", r))
}
