package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| API DATES |---------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

// Fonction pour récupérer les dates depuis l'API et les intégrer dans la structure globale
func Api_dates() []Globale_Structure_Donnees {
	dat := "https://groupietrackers.herokuapp.com/api/dates"

	// Faire une requête HTTP GET
	resp, err := http.Get(dat)
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

	var apiDates ApiDATES
	if err := json.Unmarshal(body, &apiDates); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

	// Si la liste des dates est vide, retourner une liste vide
	if len(apiDates.Index) == 0 {
		return nil
	}

	// Créer une liste pour stocker les informations de toutes les dates
	var allDates []Globale_Structure_Donnees

	// Parcourir chaque index et remplir la structure Globale_Structure_Donnees
	for _, date := range apiDates.Index {
		// Créer une nouvelle instance de Globale_Structure_Donnees pour chaque date
		dateData := Globale_Structure_Donnees{
			ID_DATES:    date.Id,
			DATES_DATES: date.Dates,
			INDEX_DATES: len(date.Dates),
		}

		// Ajouter l'objet dateData à la liste allDates
		allDates = append(allDates, dateData)
	}

	// Retourner la liste de toutes les dates
	return allDates
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
