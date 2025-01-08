package main

import (
	"fmt"
	"groupie-tracker/opti"
	"groupie-tracker/structures"
)

func main() {

	// // Appeler les fonctions API pour récupérer les données
	// api.Api_artists()   // Remplit Apiartistes
	// api.Api_locations() // Remplit apiResponse
	// api.Api_dates()     // Remplit api.Api_const
	// api.Api_DL()        // Remplit api_dl

	opti.DONNEES_API()

	Affichage_Apres_Triage_Dates()
}

func Affichage_Apres_Triage_Dates() {

	// Vérifier si api.Api_const.Index contient des données
	if len(structures.IndexDates.Index) > 0 {
		// Parcours chaque "Dates" dans "api.Api_const.Index"
		for i, DATES := range structures.IndexDates.Index {
			// Appeler la fonction pour traiter les dates
			Dates, _ := opti.Enlever_Caracteres_Dates(DATES.Dates, "")
			DatesF := opti.Triage_Date(Dates)

			// Afficher les dates traitées
			fmt.Println("Id = ", i)
			for j, date := range DatesF {
				fmt.Println(j, "   ", DATES.Dates[j], " ou ", Dates[j], " ou ", date)
			}
		}
	} else {
		fmt.Println("Aucune date disponible dans api.Api_const.")
	}

	// Ligne de séparation
	fmt.Println("\n", "--------------------------------------------------------------------------------------------", "\n")

	// Parcours chaque artiste dans "api.Apiartistes"
	for i, artiste := range structures.Apiartistes {
		// Convertir artiste.FirstAlbum en slice de chaînes ([]string)
		FirstAlbumSlice := []string{artiste.FirstAlbum}

		// Appeler la fonction pour traiter les dates de l'artiste
		First, _ := opti.Enlever_Caracteres_Dates(FirstAlbumSlice, "")
		AFirst := opti.Triage_Date(First)

		// Afficher les dates traitées
		for j, first := range AFirst {
			fmt.Println("Id = ", i, "  ", artiste.FirstAlbum, " ou ", First[j], " ou ", first)
		}
	}
}
