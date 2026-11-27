package structure_donnee

type Equipe struct {
	Id                  int    `json:"id"`
	Equipe_lien			[]string `json:"pokemons"`
	Equipe_pokemon		[]Pokemon 
}