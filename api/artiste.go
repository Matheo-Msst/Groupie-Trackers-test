package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| API ARTISTES |------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

// Fonction pour récupérer tous les artistes depuis l'API
func Api_artists() []Globale_Structure_Donnees {
	// URL de l'API
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

	// Décode le JSON dans la variable ApiArtistes
	var ApiArtistes []Artiste
	if err := json.Unmarshal(body, &ApiArtistes); err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

	// Si la liste des artistes est vide, retourner une liste vide
	if len(ApiArtistes) == 0 {
		return nil
	}
	// Créer une liste pour stocker les informations de tous les artistes
	var AllArtists []Globale_Structure_Donnees

	// Parcourir chaque artiste et remplir la structure Globale_Structure_Donnees
	for _, artiste := range ApiArtistes {
		// Créer un objet Globale_Structure_Donnees pour chaque artiste
		artistData := Globale_Structure_Donnees{
			ID_ARTISTE:            artiste.Id,
			IMAGE:                 artiste.Image,
			NAME_ARTISTE:          artiste.Name,
			MEMBERS_ARTISTE:       artiste.Members,
			CREATION_DATE_ARTISTE: artiste.CreationDate,
			FIRST_ALBUM_ARTISTE:   artiste.FirstAlbum,
			LOCATIONS_ARTISTE:     artiste.Locations,
			CONCERT_DATES_ARTISTE: artiste.ConcertDates,
			RELATIONS_ARTISTE:     artiste.Relations,
		}

		// Ajouter l'artiste à la liste allArtists
		AllArtists = append(AllArtists, artistData)
	}

	// Retourner la liste de tous les artistes
	return AllArtists
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
