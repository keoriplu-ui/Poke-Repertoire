package main

import (
	"fmt"
	"poke-repertoire/scr/golang/recuperation"
)

func main() {
	poke, err := recuperation.RecuperationPokemon()
	if err != "" {
		fmt.Println(err)
	}
	dress , err := recuperation.RecuperationDresseur()
	if err != "" {
		fmt.Println(err)
	}
	fmt.Println(poke, dress)
}
