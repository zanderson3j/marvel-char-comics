package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) character(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")
	characterList := app.marvelClient.DisplayCharacterList(name)
	json.NewEncoder(w).Encode(characterList)
}

func (app *application) comics(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	w.Header().Set("Content-Type", "application/json")
	characterList := app.marvelClient.DisplayComicsAndCharactersForCharacterID(id)
	json.NewEncoder(w).Encode(characterList)
}
