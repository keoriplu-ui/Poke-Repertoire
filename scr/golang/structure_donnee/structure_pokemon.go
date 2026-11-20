package structure_donnee

type Pokemon struct {
	Id 					int `json:"id"`
	Nom 				string `json:"name"`
	Image 				string `json:"image"`
	Description 		string `json:"string"`
	Premier_type 		string `json:"first_type"`
	Second_type 		string `json:"second_type"`
	Generation 			int `json:"generation"`
	Attaque_signature 	string `json:"signature_attack"`
}