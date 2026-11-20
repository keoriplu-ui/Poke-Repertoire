package structure_donnee

type Dresseur struct {
	Id                  int    `json:"id"`
	Nom                 string `json:"name"`
	Age                 int `json:"age"`
	Region_origine      string `json:"origin_region"`
	Ville_origine       string `json:"origin_city"`
	Meilleur_partenaire string `json:"best_buddy"`
	Equipe              string `json:"team"`
}