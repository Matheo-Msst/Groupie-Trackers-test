package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Api_dates() {
	dat := "https://groupietrackers.herokuapp.com/api/dates"

	// Faire une requête HTTP GET
	resp, err := http.Get(dat)
	if err != nil {
		fmt.Printf("Erreur lors de la requête HTTP GET : %v", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Printf("Erreur lors de la fermeture du corps de la réponse : %v", cerr)
		}
	}()

	// Lire la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture de la réponse : %v", err)
	}

	// Désérialiser les données JSON
	if err := json.Unmarshal(body, &IndexDates); err != nil {
		fmt.Printf("Erreur lors du décodage du JSON : %v", err)
	}

}
