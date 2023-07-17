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

	fsAssets := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fsAssets))
	fsScripts := http.FileServer(http.Dir("scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fsScripts))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/artists", handlers.ArtistList)
	http.HandleFunc("/artist/", handlers.ArtistInfos)
	http.HandleFunc("/events", handlers.DatesList)
	http.HandleFunc("/suggestion", handlers.Suggestion)
	http.HandleFunc("/search", handlers.Search)
	http.HandleFunc("/map/", handlers.Map)

	log.Println("Server started and listening on", PORT)
	log.Println(ADDRESS + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
