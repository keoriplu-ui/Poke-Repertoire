package recuperation_pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poke-repertoire/src/golang/structure_donnee"
	"time"
)

func Recuperation_pokemon(lien string) ([]structure_donnee.Pokemon, string) {
	client := http.Client{Timeout: 60 * time.Second}
	reponse, err := client.Get("https://poke-repertoire.vercel.app/api/teams")
	if err != nil {
		fmt.Println(err)
		return []structure_donnee.Pokemon{}, err.Error()
	}
	if reponse.StatusCode != 200 {
		fmt.Println(err)
		reponse.Body.Close()
		return []structure_donnee.Pokemon{}, reponse.Status
	}
	var pokemons []structure_donnee.Pokemon
	err = json.NewDecoder(reponse.Body).Decode(&pokemons)
	reponse.Body.Close()
	if err != nil {
		fmt.Println(err)
		return []structure_donnee.Pokemon{}, err.Error()
	}
	return pokemons, ""
}
