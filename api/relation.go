package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| API RELATIONS |-----------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

// Fonction pour récupérer les relations depuis l'API et les intégrer dans la structure globale
func Api_relations() []Globale_Structure_Donnees {
	// URL de l'API pour récupérer les relations
	dl := "https://groupietrackers.herokuapp.com/api/relation"

	// Faire une requête HTTP GET
	resp, err := http.Get(dl)
	if err != nil {
		log.Fatalf("Erreur lors de la requête HTTP GET : %v", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Printf("Erreur lors de la fermeture du corps de la réponse : %v", cerr)
		}
	}()

	// Lire la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse : %v", err)
	}

	// Désérialiser le JSON dans la structure Api_Relations
	var apiRelations Api_Relations
	if err := json.Unmarshal(body, &apiRelations); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

	// Si la liste des relations est vide, retourner une liste vide
	if len(apiRelations.Index) == 0 {
		return nil
	}

	// Créer une liste pour stocker les informations de toutes les relations
	var allRelations []Globale_Structure_Donnees

	// Parcourir chaque relation et remplir la structure Globale_Structure_Relations
	for _, relation := range apiRelations.Index {
		var locations []string
		var dates []string
		datesLocationMap := make(map[string][]string)

		// Itérer à travers la map des lieux et dates
		for lieu, dateList := range relation.DatesLocalisation {
			locations = append(locations, lieu) // Ajouter le lieu
			datesLocationMap[lieu] = dateList   // Ajouter les dates pour ce lieu
		}

		relationData := Globale_Structure_Donnees{
			ID_RELATIONS:             relation.Id,
			LOCATIONS_RELATIONS:      locations,
			DATES_LOCATIONS:          dates,
			INDEX_RELATIONS:          len(relation.DatesLocalisation),
			DATESLOCATIONS_RELATIONS: datesLocationMap,
		}

		// Ajouter l'objet relationData à la liste allRelations
		allRelations = append(allRelations, relationData)
	}

	// Retourner la liste de toutes les relations
	return allRelations
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
