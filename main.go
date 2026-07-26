package main

import (
	"fmt"
	"log"
	"lisen-words/pkg/deezer"
)

func main() {

	res, err := deezer.SearchTrack("bullet for my valentine bittersweet ")
	if err != nil {
		log.Fatalf("Error al buscar: %v", err)
	}

	for i, track := range res.Data {
		fmt.Printf("[%d] %s - %s (ID: %d)\n", i+1, track.Title, track.Artist.Name, track.ID)
		fmt.Println("----------------------------------------")
	}
}