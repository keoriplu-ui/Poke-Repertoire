package recuperation_equipe

import (
	"fmt"
	recuperation_equipe "poke-repertoire/src/golang/recuperation/recuperation_equipe/recuperation_separe"
	"poke-repertoire/src/golang/structure_donnee"
)

func RecuperationEquipe() ([]structure_donnee.Equipe, string) {
	equipes, err := recuperation_equipe.RecuperationEquipeLien()
	if err != "" {
		fmt.Println(err)
		return []structure_donnee.Equipe{}, err
	}
	equipes, err = recuperation_equipe.RecuperationEquipePokemon(equipes)
	if err != "" {
		fmt.Println(err)
		return []structure_donnee.Equipe{}, err
	}
	return equipes, ""
}
