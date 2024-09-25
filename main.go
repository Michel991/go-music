package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Music struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Artist *Artist `json:"director"`
}

type Artist struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musics)

}

func deleteMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range musics {
		if item.ID == params["id"] {
			musics = append(musics[:index], musics[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(musics)
	http.Error(w, "Music not found", http.StatusNotFound)
}

func getMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range musics {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func updateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range musics {
		if item.ID == params["id"] {
			musics = append(musics[:index], musics[index+1:]...)
			var music Music
			_ = json.NewDecoder(r.Body).Decode(&music)
			music.ID = params["id"]
			musics = append(musics, music)
			json.NewEncoder(w).Encode(music)

		}

	}

}

func createMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var music Music
	_ = json.NewDecoder(r.Body).Decode(&music)
	music.ID = strconv.Itoa(rand.Intn(100000))
	musics = append(musics, music)
	json.NewEncoder(w).Encode(music)

}

var musics []Music

func main() {

	r := mux.NewRouter()

	musics = append(musics, Music{ID: "1", Isbn: "4324", Title: "Terminator", Artist: &Artist{Firstname: "Michel", Lastname: "Balla"}})
	musics = append(musics, Music{ID: "2", Isbn: "5324", Title: "ShoeDog", Artist: &Artist{Firstname: "Phil", Lastname: "Knight"}})
	musics = append(musics, Music{ID: "3", Isbn: "6324", Title: "The last Don", Artist: &Artist{Firstname: "Mario", Lastname: "Puzzo"}})
	musics = append(musics, Music{ID: "4", Isbn: "7324", Title: "Prince Of Fire", Artist: &Artist{Firstname: "Daniel", Lastname: "Silva"}})

	r.HandleFunc("/musics", getMusics).Methods("GET")
	r.HandleFunc("/musics/{id}", getMusic).Methods("GET")
	r.HandleFunc("/musics", createMusic).Methods("POST")
	r.HandleFunc("/musics/{id}", updateMusic).Methods("PUT")
	r.HandleFunc("/musics/{id}", deleteMusic).Methods("DELETE")

	fmt.Fprintln(os.Stdout, "Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
