package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`	
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	
	firstName string `json:"firstName"`
	lastName string `json:"lastName"`

}

var movie []Movie

func main(
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "The Shawshank Redemption", Director: &Director{firstName: "Frank", lastName: "Darabont"}})
	movies = append(movies, Movie{ID: "2", Isbn: "432532", Title: "Godzilla", Director: &Director{firstName: "Steven", lastName: "Spielberg"}})
	r.HandleFunc("/movie", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
)