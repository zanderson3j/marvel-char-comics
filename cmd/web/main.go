package main

import (
	"log"
	"net/http"

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

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
