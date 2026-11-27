package recuperation_equipe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poke-repertoire/src/golang/structure_donnee"
	"time"
)

func RecuperationEquipeLien() ([]structure_donnee.Equipe, string) {
	client := http.Client{Timeout: 60 * time.Second}
	reponse, err := client.Get("https://poke-repertoire.vercel.app/api/teams")
	if err != nil {
		fmt.Println(err)
		return []structure_donnee.Equipe{}, err.Error()
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != 200 {
		fmt.Println(err)
		return []structure_donnee.Equipe{}, reponse.Status
	}
	var equipes []structure_donnee.Equipe
	err = json.NewDecoder(reponse.Body).Decode(&equipes)
	if err != nil {
		fmt.Println(err)
		return []structure_donnee.Equipe{}, err.Error()
	}
	return equipes, ""
}
