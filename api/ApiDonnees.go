package api

import (
	"encoding/json"
	"groupie-tracker/structures"
	"io"
	"log"
	"net/http"
)

func RecupDonneesApi() {
	URL_ARTISTES := "https://groupietrackers.herokuapp.com/api/artists"    //----------------------------------URL ARTISTES
	URL_DATES := "https://groupietrackers.herokuapp.com/api/dates"         //----------------------------------URL DATES
	URL_LOCATIONS := "https://groupietrackers.herokuapp.com/api/locations" //----------------------------------URL LOCATIONS
	URL_RELATIONS := "https://groupietrackers.herokuapp.com/api/relation"  //----------------------------------URL RELATIONS

	// ------------------------------------------------ REQUETE HTTP GET -----------------------------------------------

	// Faire une requête HTTP GET ------------------------------------------------------------------------ARTISTES
	resp_Artistes, err_Artistes := http.Get(URL_ARTISTES)
	if err_Artistes != nil {
		log.Fatalf("Erreur lors de la requête HTTP GET : %v", err_Artistes)
	}
	defer func() {
		if cerr := resp_Artistes.Body.Close(); cerr != nil {
			log.Printf("Erreur lors de la fermeture du corps de la réponse : %v", cerr)
		}
	}()
	// Faire une requête HTTP GET ------------------------------------------------------------------------DATES
	resp_Dates, err_Dates := http.Get(URL_DATES)
	if err_Dates != nil {
		log.Fatalf("Erreur lors de la requête HTTP GET : %v", err_Dates)
	}
	defer func() {
		if cerr := resp_Dates.Body.Close(); cerr != nil {
			log.Printf("Erreur lors de la fermeture du corps de la réponse : %v", cerr)
		}
	}()
	// Faire une requête HTTP GET ------------------------------------------------------------------------LOCATIONS
	resp_Locations, err_Locations := http.Get(URL_LOCATIONS)
	if err_Locations != nil {
		log.Fatalf("Erreur lors de la requête HTTP GET : %v", err_Locations)
	}
	defer func() {
		if cerr := resp_Locations.Body.Close(); cerr != nil {
			log.Printf("Erreur lors de la fermeture du corps de la réponse : %v", cerr)
		}
	}()
	// Faire une requête HTTP GET
	resp_Relations, err_Relations := http.Get(URL_RELATIONS)
	if err_Relations != nil {
		log.Fatalf("Erreur lors de la requête HTTP GET : %v", err_Relations)
	}
	defer func() {
		if cerr := resp_Relations.Body.Close(); cerr != nil {
			log.Printf("Erreur lors de la fermeture du corps de la réponse : %v", cerr)
		}
	}()

	// ------------------------------------------------ LIRE LA REPONSE (body) -----------------------------------------------

	// Lire la réponse ------------------------------------------------------------------------ARTISTES
	body_Artistes, err_Artistes := io.ReadAll(resp_Artistes.Body)
	if err_Artistes != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse : %v", err_Artistes)
	}
	// Lire la réponse ------------------------------------------------------------------------DATES
	body_Dates, err_Dates := io.ReadAll(resp_Dates.Body)
	if err_Dates != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse : %v", err_Dates)
	}
	// Lire la réponse ------------------------------------------------------------------------LOCATIONS
	body_Locations, err_Locations := io.ReadAll(resp_Locations.Body)
	if err_Locations != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse : %v", err_Locations)
	}
	// Lire la réponse ------------------------------------------------------------------------RELATIONS
	body_Relations, err_Relations := io.ReadAll(resp_Relations.Body)
	if err_Relations != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse : %v", err_Relations)
	}

	// ------------------------------------------------ METTRE LES DONNEES DANS LES STRUCT ----------------------------------

	// Désérialiser les données JSON ------------------------------------------------------------------------ARTISTES
	if err_Artistes := json.Unmarshal(body_Artistes, &structures.Apiartistes); err_Artistes != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err_Artistes)
	}
	// Désérialiser les données JSON ------------------------------------------------------------------------DATES
	if err_Dates := json.Unmarshal(body_Dates, &structures.IndexDates); err_Dates != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err_Dates)
	}
	// Désérialiser les données JSON ------------------------------------------------------------------------LOCATIONS
	if err_Locations := json.Unmarshal(body_Locations, &structures.IndexLocations); err_Locations != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err_Locations)
	}
	// Désérialiser les données JSON ------------------------------------------------------------------------RELATIONS
	if err_Relations := json.Unmarshal(body_Relations, &structures.IndexRelations); err_Relations != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err_Relations)
	}

}

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
