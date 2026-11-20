package recuperation

import (
	"encoding/json"
	"net/http"
	"poke-repertoire/scr/golang/structure_donnee"
	"time"
)

func RecuperationDresseur() ([]structure_donnee.Dresseur, string) {
	client := http.Client{Timeout: 60 * time.Second}
	reponse, err := client.Get("https://poke-repertoire.vercel.app/api/trainers")
	if err != nil {
		return []structure_donnee.Dresseur{}, err.Error()
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != 200 {
		return []structure_donnee.Dresseur{}, reponse.Status
	}
	var dresseurs []structure_donnee.Dresseur
	err = json.NewDecoder(reponse.Body).Decode(&dresseurs)
	if err != nil {
		return []structure_donnee.Dresseur{}, err.Error()
	}
	return dresseurs, ""
}
