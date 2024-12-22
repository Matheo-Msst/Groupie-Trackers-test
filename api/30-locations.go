package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Api_Locations() {
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

	if err := json.Unmarshal(body, &IndexLocations); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

}

func Afficher_Locations() {
	for i, location := range IndexLocations.Locations {
		/* Pour afficher uniquement le nombre d'elements souhaites */
		if i >= ElementAfficher {
			break
		}

		// Affiche l'ID
		fmt.Println("\n", "ID: ", location.Id)

		// Affiche chaque lieu dans Locations
		fmt.Println("Locations :")
		for j, lieu := range location.Locations {
			fmt.Println(j, "  .  ", lieu)
		}
	}
}
