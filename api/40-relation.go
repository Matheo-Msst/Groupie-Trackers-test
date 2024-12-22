package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Api_Relations() {
	dl := "https://groupietrackers.herokuapp.com/api/relation"

	// Faire une requête HTTP GET
	resp, err := http.Get(dl)
	if err != nil {
		fmt.Printf("Erreur lors de la requête HTTP GET : %v", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			fmt.Printf("Erreur lors de la fermeture du corps de la réponse : %v", cerr)
		}
	}()

	// Lire la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture de la réponse : %v", err)
	}

	// Désérialiser les données JSON
	if err := json.Unmarshal(body, &IndexRelations); err != nil {
		fmt.Printf("Erreur lors du décodage du JSON : %v", err)
	}

}

func Afficher_Relations() {
	for /*_*/ i, relation := range IndexRelations.Relations {
		/* Pour afficher uniquement le nombre d'elements souhaites */
		if i >= ElementAfficher {
			break
		}

		// Affiche l'ID
		fmt.Printf("ID: %d\n", relation.Id)

		// Affiche chaque lieu dans DatesLocations
		fmt.Println("Lieux et dates :")
		for lieu, dates := range relation.DatesLocalisation {
			fmt.Printf("  Lieu: %s\n", lieu)
			for _, date := range dates {
				fmt.Println("    Date: ", date)
			}
		}
	}
}
