package recuperation_equipe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poke-repertoire/src/golang/structure_donnee"
	"time"
)

func RecuperationEquipePokemon(equipes []structure_donnee.Equipe) ([]structure_donnee.Equipe, string) {
	client := http.Client{Timeout: 60 * time.Second}
	for _, equipe := range equipes {
		for _, lien := range equipe.Equipe_lien {
			reponse, err := client.Get(lien)
			if err != nil {
				fmt.Println(err)
				return []structure_donnee.Equipe{}, err.Error()
			}
			if reponse.StatusCode != 200 {
				fmt.Println(err)
				reponse.Body.Close()
				return []structure_donnee.Equipe{}, reponse.Status
			}
			var pokemon []structure_donnee.Pokemon
			err = json.NewDecoder(reponse.Body).Decode(&pokemon)
			reponse.Body.Close()
			if err != nil {
				fmt.Println(err)
				return []structure_donnee.Equipe{}, err.Error()
			}
			equipe.Equipe_pokemon = append(equipe.Equipe_pokemon, pokemon...)
		}
		equipes = append(equipes, structure_donnee.Equipe{Id: equipe.Id , Equipe_lien: equipe.Equipe_lien , Equipe_pokemon:  equipe.Equipe_pokemon})
	}
	return equipes, ""
}
