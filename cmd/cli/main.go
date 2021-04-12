package main

import (
	"fmt"
	"os"

	"github.com/zanderson3j/marvel-char-comics/pkg/marvelclient"
)

func main() {
	client := &marvelclient.Marvelclient{}

	arg := os.Args[1]

	if arg == "list" {
		result := client.DisplayCharacterList(os.Args[2])
		fmt.Println(result)
	} else {
		result := client.DisplayComicsAndCharactersForCharacterID(arg)
		fmt.Println(result)
	}
}
