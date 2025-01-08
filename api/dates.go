package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"io"
	"log"
	"net/http"
)

func Api_dates() {
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

	// Désérialiser les données JSON

	if err := json.Unmarshal(body, &structures.IndexDates); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

}

func AfficherDates() {
	// Afficher les premiers éléments

	for i, dates := range structures.IndexDates.Index {
		if i >= structures.ElementAfficher {
			break
		}

		// Affiche l'ID
		fmt.Printf("ID: %d\n", dates.Id)

		// Affiche chaque lieu dans Locations
		fmt.Println("Dates :")
		for j, lieu := range dates.Dates {
			fmt.Printf("  %d. %s\n", j+1, lieu)
		}
	}
}
