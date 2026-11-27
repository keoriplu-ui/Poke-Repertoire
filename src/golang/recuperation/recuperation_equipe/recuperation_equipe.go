package recuperation_equipe

import (
	"fmt"
	"poke-repertoire/src/golang/recuperation/recuperation_equipe/recuperation_separe"
	"poke-repertoire/src/golang/structure_donnee"
)

func RecuperationEquipe() ([]structure_donnee.Equipe, string) {
	equipes, err := recuperation_separe.RecuperationEquipeLien()
	if err != "" {
		fmt.Println(err)
		return []structure_donnee.Equipe{}, err
	}
	equipes, err = recuperation_separe.RecuperationEquipePokemon(equipes)
	if err != "" {
		fmt.Println(err)
		return []structure_donnee.Equipe{}, err
	}
	return equipes, ""
}
