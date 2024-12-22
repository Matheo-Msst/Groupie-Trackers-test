package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Apiartistes []Artiste

func Api_Artists() {
	art := "https://groupietrackers.herokuapp.com/api/artists"

	// Faire une requête HTTP GET
	resp, err := http.Get(art)
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
	if err := json.Unmarshal(body, &Apiartistes); err != nil {
		fmt.Printf("Erreur lors du décodage du JSON : %v", err)
	}

}

func Afficher_Artistes() {
	for /*_*/ i, artiste := range Apiartistes {
		/* Pour afficher uniquement le nombre d'elements souhaites */
		if i >= ElementAfficher {
			break
		}
		// Affiche l'ID et le nom de l'artiste
		fmt.Println("\n", "ID: ", artiste.Id, "\n", "Nom: ", artiste.Name, "\n")
		fmt.Println("Membres :")
		// Affiche chaque membre du groupe
		for j, membre := range artiste.Members {
			fmt.Printf("  %d. %s\n", j+1, membre)
		}
	}
}
