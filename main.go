package main

import (
	"fmt"
	handlers "groupietracker/lib/handlers"
	utils "groupietracker/lib/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	utils.LoadEnv(".env")

	port := os.Getenv("PORT")
	PORT := fmt.Sprintf(":%v", port)

	ADDRESS := os.Getenv("ADDRESS")

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/artists", handlers.ArtistList)
	http.HandleFunc("/artist/", handlers.ArtistInfos)
	http.HandleFunc("/events", handlers.DatesList)
	http.HandleFunc("/suggestion", handlers.Suggestion)
	http.HandleFunc("/search", handlers.Search)

	log.Println("Server started and listening on", PORT)
	log.Println(ADDRESS + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
