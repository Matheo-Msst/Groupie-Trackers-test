package opti

import (
	"encoding/json"
	"groupie-tracker/structures"
	"io"
	"log"
	"net/http"
)

func DONNEES_API() {
	// les URLs et les pointeurs vers les structures
	urlApi := []string{"artists", "dates", "locations", "relation"}
	structApi := []interface{}{&structures.Apiartistes, &structures.IndexDates, &structures.IndexLocations, &structures.IndexRelations}

	for i, urlPartie2 := range urlApi {
		// CONSTRUIRE L'URL COMPLETE
		URL := "https://groupietrackers.herokuapp.com/api/" + urlPartie2

		// REQUETE HTTP GET AVEC L'URL
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatalf("Erreur lors de la requête HTTP GET : %v", err)
		}
		defer func() {
			if cerr := resp.Body.Close(); cerr != nil {
				log.Printf("Erreur lors de la fermeture du corps de la réponse : %v", cerr)
			}
		}()

		// LIRE LA REPONSE (body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Erreur lors de la lecture de la réponse : %v", err)
		}

		// METTRE LES DONNEES DANS LES STRUCT
		if err := json.Unmarshal(body, structApi[i]); err != nil {
			log.Fatalf("Erreur lors du décodage du JSON : %v", err)
		}
	}
}
