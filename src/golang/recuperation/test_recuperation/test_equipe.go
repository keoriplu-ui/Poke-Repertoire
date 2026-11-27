package test_recuperation

import (
	"fmt"
	"poke-repertoire/src/golang/recuperation/recuperation_equipe"
)

func TestRecuperationEquipe() {
	equipes, err := recuperation_equipe.RecuperationEquipeLien()
	if err != "" {
		fmt.Println(err)
		return
	}
	for _, equipe := range equipes {
		fmt.Println(equipe.Equipe_pokemon)
	}
}
