package test_recuperation

import (
	"fmt"
	"poke-repertoire/src/golang/recuperation/recuperation_equipe"
)

func TestRecuperationEquipe() {
	equipes, err := recuperation_equipe.RecuperationEquipe()
	if err != "" {
		fmt.Println(err)
		fmt.Println("a")
		return
	}
	for _, equipe := range equipes {
		fmt.Println(equipe.Equipe_pokemon)
	}
}
