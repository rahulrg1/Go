package main

import (
	"fmt"
	"go-psql-crud/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// type Movie struct {
// 	ID       string    `json:"id"`
// 	Isbn     string    `json:"isbn"`
// 	Title    string    `json:"title"`
// 	Director *Director `json:"director"`
// }

// type Director struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// var movies []Movie

// func getMovies(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// params:=
// 	// fmt.Println(params)
// 	json.NewEncoder(w).Encode(movies)
// }

// func deleteMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[0:index], movies[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(movies)
// }

// func getMovie(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	for _, item := range movies {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}

// 	}
// }

// func createMovie(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	var movie Movie
// 	_= json.NewDecoder(r.Body).Decode(&movie)
// 	fmt.Println(r.Body)
// 	movie.ID = strconv.Itoa(rand.Intn(1000000))
// 	movies = append(movies, movie)
// 	json.NewEncoder(w).Encode(movie)
// }

// func updateMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			var movie Movie
// 			_ = json.NewDecoder(r.Body).Decode(&movie)
// 			movie.ID = params["id"]
// 			movies = append(movies, movie)
// 			json.NewEncoder(w).Encode(movie)
// 			return
// 		}
// 	}

// }

// func main() {
// 	r := mux.NewRouter()
// 	var mov Movie
// 	fmt.Println(mov)
// 	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One",
// 		Director: &Director{Firstname: "run", Lastname: "roy"}})
// 	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two",
// 		Director: &Director{Firstname: "rahul", Lastname: "Garai"}})
// 	r.HandleFunc("/movies", getMovies).Methods("GET")
// 	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
// 	r.HandleFunc("/movies", createMovie).Methods("POST")
// 	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
// 	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

// 	fmt.Printf("server started PORT 8000!..")
// 	log.Fatal(http.ListenAndServe(":8000", r))
// }
