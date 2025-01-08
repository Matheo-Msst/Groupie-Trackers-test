package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"io"
	"log"
	"net/http"
)

func Api_artists() {
	art := "https://groupietrackers.herokuapp.com/api/artists"

	// Faire une requête HTTP GET
	resp, err := http.Get(art)
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

	if err := json.Unmarshal(body, &structures.Apiartistes); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

}

func AfficherArtistes() {
	// Afficher les premiers éléments

	for i, artiste := range structures.Apiartistes {
		if i >= structures.ElementAfficher {
			break
		}
		// Affiche l'ID et le nom de l'artiste
		fmt.Printf("ID: %d | Nom: %s\n", artiste.Id, artiste.Name)
		fmt.Println("Membres :")
		// Affiche chaque membre du groupe
		for j, membre := range artiste.Members {
			fmt.Printf("  %d. %s\n", j+1, membre)
		}
	}
}
