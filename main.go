package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"math/rand"
	"net/http"
	"log"
	"strconv"
)

type Movie struct{
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`
}

type Director struct{
	FirstName string  `json: "firstname"`
	LastName string  `json: "lastname"`
}

var movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.Encoder(w).Encoder(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1]...)
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range moview {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request ){
	w.Header().Add("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

// UpdateMovie function

func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "4321", Title: "El mañana", Director : &Director{FirstName: "Vinicio",LastName: "Del Toro"})
	movies = append(movies, Movie{ID: "2", Isbn: "432r", Title: "El ayer", Director : &Director{FirstName: "Vinicio",LastName: "Del Mono"})


	router := mux.NewRouter()

	router.HandleFunc("/movies", GetMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", GetMovie).Methods("GET")
	router.HandleFunc("/movies", CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting port: 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
