package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"io"
	"log"
	"net/http"
)

func Api_DL() {
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

	// Désérialiser les données JSON

	if err := json.Unmarshal(body, &structures.IndexRelations); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

}

func AfficherRelations() {
	// Afficher les premiers éléments

	for i, relation := range structures.IndexRelations.Index {
		if i >= structures.ElementAfficher {
			break
		}

		// Affiche l'ID
		fmt.Printf("ID: %d\n", relation.Id)

		// Affiche chaque lieu dans DatesLocations
		fmt.Println("Lieux et dates :")
		for lieu, dates := range relation.DatesLocalisation {
			fmt.Printf("  Lieu: %s\n", lieu)
			for _, date := range dates {
				fmt.Printf("    Date: %s\n", date)
			}
		}
	}
}
