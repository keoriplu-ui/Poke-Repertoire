package structure_donnee

type Equipe struct {
	Id                  int    `json:"id"`
	Equipe_lien			string `json:"pokemon"`
	Equipe_pokemon		[]Pokemon 
}