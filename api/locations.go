package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| API LOCATIONS |-----------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

// Fonction pour récupérer les données de localisation depuis l'API
func Api_locations() []Globale_Structure_Donnees {
	loc := "https://groupietrackers.herokuapp.com/api/locations"

	// Faire une requête HTTP GET
	resp, err := http.Get(loc)
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

	// Désérialiser le JSON dans la structure ApiResponse
	if err := json.Unmarshal(body, &ApiLocations); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

	// Si la liste des localisations est vide, retourner une liste vide
	if len(ApiLocations.Index) == 0 {
		return nil
	}

	// Créer une liste pour stocker les informations de toutes les localisations
	var allLocations []Globale_Structure_Donnees

	// Parcourir chaque localisation et remplir la structure Globale_Structure_Donnees
	for _, location := range ApiLocations.Index {
		// Créer une nouvelle instance de Globale_Structure_Donnees pour chaque localisation
		locationData := Globale_Structure_Donnees{
			ID_LOCATIONS:        location.Id,
			LOCATIONS_LOCATIONS: location.Locations,
			DATES_LOCATIONS:     []string{location.Dates},
		}

		// Ajouter l'objet locationData à la liste allLocations
		allLocations = append(allLocations, locationData)

	}

	// Retourner la liste de toutes les localisations
	return allLocations
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
