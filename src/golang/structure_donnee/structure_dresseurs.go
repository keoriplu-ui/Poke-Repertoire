package structure_donnee

type Dresseur struct {
	Id               int     `json:"id"`
	Nom              string  `json:"name"`
	Age              int     `json:"age"`
	Region_d_origine string  `json:"origin_region"`
	Ville_d_origine  string  `json:"origin_city"`
	Pokemon_favori   Pokemon `json:"best_buddy"`
	Equipe           Equipe  `json:"team"`
}
