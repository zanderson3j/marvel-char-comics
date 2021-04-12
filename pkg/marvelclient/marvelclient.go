package marvelclient

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/zanderson3j/marvel-char-comics/pkg/models"
)

// Marvel API Keys
var (
	pubkey  = os.Getenv("PUBKEY")
	privkey = os.Getenv("PRIVKEY")
)

type Marvelclient struct {
}

func hashCalculation(ts string) string {
	hasher := md5.New()
	hasher.Write([]byte(ts + privkey + pubkey))

	return hex.EncodeToString(hasher.Sum(nil))
}

func callMarvel(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func (mc *Marvelclient) DisplayComicsAndCharactersForCharacterID(characterID string) *models.MarvelComicResponse {
	ts := strconv.Itoa(time.Now().Second())
	hash := hashCalculation(ts)
	url := "https://gateway.marvel.com:443/v1/public/characters/" + characterID + "/comics?orderBy=focDate" + "&ts=" + ts + "&apikey=" + pubkey + "&hash=" + hash

	responseBody := callMarvel(url)

	var comicDataWrapper models.ComicDataWrapper
	if err := json.Unmarshal(responseBody, &comicDataWrapper); err != nil {
		panic(err)
	}

	resultsLength := comicDataWrapper.Data.Count
	comicsSlice := make([]models.MarvelComic, resultsLength)

	for i, comic := range comicDataWrapper.Data.Results {
		characterSlice := make([]string, comic.Characters.Returned)

		for j, character := range comic.Characters.Items {
			characterSlice[j] = character.Name
		}

		comicsSlice[i] = models.MarvelComic{Title: comic.Title, Characters: characterSlice}
	}

	return &models.MarvelComicResponse{ResultsCount: resultsLength, Comics: comicsSlice}
}

func (mc *Marvelclient) DisplayCharacterList(name string) *models.MarvelCharacterResponse {
	ts := strconv.Itoa(time.Now().Second())
	hash := hashCalculation(ts)
	url := "https://gateway.marvel.com:443/v1/public/characters?nameStartsWith=" + name + "&ts=" + ts + "&apikey=" + pubkey + "&hash=" + hash

	responseBody := callMarvel(url)

	var characterDataWrapper models.CharacterDataWrapper
	if err := json.Unmarshal(responseBody, &characterDataWrapper); err != nil {
		panic(err)
	}

	resultLength := characterDataWrapper.Data.Count
	characterSlice := make([]models.MarvelCharacter, resultLength)

	for i, character := range characterDataWrapper.Data.Results {
		characterSlice[i] = models.MarvelCharacter{Name: character.Name, CharacterID: character.ID}
	}

	return &models.MarvelCharacterResponse{ResultsCount: resultLength, Characters: characterSlice}
}
