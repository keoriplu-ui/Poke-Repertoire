package recuperation_separe

import (
	"fmt"
	"poke-repertoire/src/golang/recuperation/recuperation_equipe/recuperation_separe/recuperation_unique_pokemon"
	"poke-repertoire/src/golang/structure_donnee"
)

func RecuperationEquipePokemon(equipes []structure_donnee.Equipe) ([]structure_donnee.Equipe, string) {

	for index, equipe := range equipes {
		ajout_pokemons := []structure_donnee.Pokemon{}
		for _, lien := range equipe.Equipe_lien {
			pokemon, err := recuperation_unique_pokemon.Recuperation_unique_pokemon(lien)
			if err != "" {
				fmt.Println(err)
				continue
			}
			ajout_pokemons = append(ajout_pokemons, pokemon)
		}
		equipes[index].Equipe_pokemon = append(equipes[index].Equipe_pokemon, ajout_pokemons...)
	}
	return equipes, ""
}
