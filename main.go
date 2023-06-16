package main

import (
	handlers "groupie_tracker/lib/handlers"
	"log"
	"net/http"
)

const ADDRESS = "http://localhost"
const PORT = ":8080"

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/artists/", handlers.ArtistList)
	http.HandleFunc("/artist/", handlers.ArtistInfos)

	log.Println("Server started and listening on", PORT)
	log.Println(ADDRESS + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
