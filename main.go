package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	parms := mux.Vars(r)
	for _, item := range movies {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	parms := mux.Vars(r)
	for index, item := range movies {
		if item.ID == parms["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	parms := mux.Vars(r)
	var movie Movie
	for index, item := range movies {
		if item.ID == parms["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			movie.ID = item.ID
			json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", ISBN: "23421", Name: "Batman", Director: &Director{FirstName: "Waleed", LastName: "khan"}})
	movies = append(movies, Movie{ID: "2", ISBN: "92103", Name: "Superman", Director: &Director{FirstName: "Moaz", LastName: "Niazi"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("connected to server 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
