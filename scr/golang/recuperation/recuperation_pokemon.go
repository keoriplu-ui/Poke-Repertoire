package recuperation

import (
	"encoding/json"
	"net/http"
	"poke-repertoire/scr/golang/structure_donnee"
	"time"
)

func RecuperationPokemon() ([]structure_donnee.Pokemon, string) {
	client := http.Client{Timeout: 60 * time.Second}
	reponse, err := client.Get("https://poke-repertoire.vercel.app/api/pokemons")
	if err != nil {
		return []structure_donnee.Pokemon{}, err.Error()
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != 200 {
		return []structure_donnee.Pokemon{}, reponse.Status
	}
	var pokemons []structure_donnee.Pokemon
	err = json.NewDecoder(reponse.Body).Decode(&pokemons)
	if err != nil {
		return []structure_donnee.Pokemon{}, err.Error()
	}
	return pokemons , ""
}
