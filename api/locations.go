package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"io"
	"log"
	"net/http"
)

func Api_locations() {
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

	// Désérialiser les données JSON

	if err := json.Unmarshal(body, &structures.IndexLocations); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

}

func AfficherLieux() {
	// Afficher les premiers éléments
	for i, location := range structures.IndexLocations.Index {
		if i >= structures.ElementAfficher {
			break
		}

		// Affiche l'ID
		fmt.Printf("ID: %d\n", location.Id)

		// Affiche chaque lieu dans Locations
		fmt.Println("Locations :")
		for j, lieu := range location.Locations {
			fmt.Printf("  %d. %s\n", j+1, lieu)
		}
	}
}
