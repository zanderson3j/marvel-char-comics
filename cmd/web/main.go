package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zanderson3j/marvel-char-comics/pkg/marvelclient"
)

type application struct {
	marvelClient *marvelclient.Marvelclient
}

func main() {
	mux := http.NewServeMux()

	app := &application{
		marvelClient: &marvelclient.Marvelclient{},
	}

	mux.HandleFunc("/character", app.character)
	mux.HandleFunc("/character/comics", app.comics)

	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
