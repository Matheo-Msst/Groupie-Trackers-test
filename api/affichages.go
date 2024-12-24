package api

import "fmt"

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| AFFICHAGE ARTISTES |------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

func Afficher_Artistes(allArtists []Globale_Structure_Donnees) {

	// Affichage de toutes les informations des artistes
	fmt.Println("Tous les artistes récupérés :")
	for _, artist := range allArtists {
		fmt.Println("ID de l'artiste:", artist.ID_ARTISTE)
		fmt.Println("Image de l'artiste:", artist.IMAGE)
		fmt.Println("Nom de l'artiste:", artist.NAME_ARTISTE)
		fmt.Println("Membres de l'artiste:", artist.MEMBERS_ARTISTE)
		fmt.Println("Date de création de l'artiste:", artist.CREATION_DATE_ARTISTE)
		fmt.Println("Premier album de l'artiste:", artist.FIRST_ALBUM_ARTISTE)
		fmt.Println("Localisations de l'artiste:", artist.LOCATIONS_ARTISTE)
		fmt.Println("Dates des concerts de l'artiste:", artist.CONCERT_DATES_ARTISTE)
		fmt.Println("Relations de l'artiste:", artist.RELATIONS_ARTISTE)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

func Afficher_Premiers_ID_Artistes(allArtists []Globale_Structure_Donnees) {

	// Vérifier qu'il y a au moins deux artistes
	if len(allArtists) < ElementsAffichages {
		fmt.Println("Il n'y a pas assez d'artistes.")
		return
	}

	// Affichage des deux premiers ID des artistes
	fmt.Println("Affichage des deux premiers ID d'artistes :")
	for i := 0; i < 2; i++ {
		artist := allArtists[i]

		fmt.Printf("ID de l'artiste: %d\n", artist.ID_ARTISTE)
		// Affichage des autres informations associées pour chaque ID
		fmt.Println("Nom de l'artiste:", artist.NAME_ARTISTE)
		fmt.Println("Image de l'artiste:", artist.IMAGE)
		fmt.Println("Membres de l'artiste:", artist.MEMBERS_ARTISTE)
		fmt.Println("Date de création de l'artiste:", artist.CREATION_DATE_ARTISTE)
		fmt.Println("Premier album de l'artiste:", artist.FIRST_ALBUM_ARTISTE)
		fmt.Println("Localisations de l'artiste:", artist.LOCATIONS_ARTISTE)
		fmt.Println("Dates des concerts de l'artiste:", artist.CONCERT_DATES_ARTISTE)
		fmt.Println("Relations de l'artiste:", artist.RELATIONS_ARTISTE)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| AFFICHAGE DATES |---------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

func Afficher_Dates(allDates []Globale_Structure_Donnees) {

	// Affichage de toutes les localisations
	fmt.Println("Toutes les Dates récupérées :")
	for _, date := range allDates {
		fmt.Println("ID de la Date:", date.ID_DATES)
		fmt.Println("Dates:", date.DATES_DATES)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

func Afficher_Premiers_ID_Dates(allDates []Globale_Structure_Donnees) {

	// Vérifier qu'il y a au moins deux dates
	if len(allDates) < ElementsAffichages {
		fmt.Println("Il n'y a pas assez de dates.")
		return
	}

	// Affichage des deux premiers ID des dates
	fmt.Println("Affichage des deux premiers ID de dates :")
	for i := 0; i < 2; i++ {
		date := allDates[i]
		fmt.Printf("ID de la Date: %d\n", date.ID_DATES)
		// Affichage des dates associées pour chaque ID
		for _, dateItem := range date.DATES_DATES {
			fmt.Printf("  Date: %s\n", dateItem)
		}
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| AFFICHAGE LOCATIONS |-----------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

func Afficher_Locations(allLocations []Globale_Structure_Donnees) {

	// Affichage de toutes les localisations
	fmt.Println("Toutes les localisations récupérées :")
	for _, location := range allLocations {
		fmt.Println("ID de la localisation:", location.ID_LOCATIONS)
		fmt.Println("Localisations:", location.LOCATIONS_LOCATIONS)
		fmt.Println("Dates:", location.DATES_LOCATIONS)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

// Fonction pour afficher les deux premiers éléments des localisations
func Afficher_Premiers_ID_Locations(allLocations []Globale_Structure_Donnees) {

	// Vérifier s'il y a au moins 2 localisations
	if len(allLocations) < ElementsAffichages {
		fmt.Println("Il y a moins de deux localisations disponibles.")
		return
	}

	// Affichage des deux premiers éléments de la liste
	fmt.Println("Les deux premiers éléments des localisations récupérées :")
	for i := 0; i < 2; i++ {
		location := allLocations[i]
		fmt.Println("ID de la localisation:", location.ID_LOCATIONS)
		fmt.Println("Localisations:", location.LOCATIONS_LOCATIONS)
		fmt.Println("Dates:", location.DATES_LOCATIONS)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| AFFICHAGE RELATIONS |-----------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

func Afficher_Relations(allRelations []Globale_Structure_Donnees) {

	// Affichage de toutes les relations récupérées
	fmt.Println("Toutes les Relations récupérées :")
	for _, relation := range allRelations {
		fmt.Printf("ID de la Relation: %d\n", relation.ID_RELATIONS)

		// Affichage des lieux et des dates associées pour chaque relation
		for i, lieu := range relation.LOCATIONS_RELATIONS {
			// Affiche le lieu
			fmt.Printf("Lieu: %s\n", lieu)
			// Affiche les dates associées au lieu
			for _, date := range relation.DATESLOCATIONS_RELATIONS[lieu] {
				fmt.Printf("  Date: %s\n", date)
			}
			// Séparer les lieux
			if i < len(relation.LOCATIONS_RELATIONS)-1 {
				fmt.Println("------------------------------------------------")
			}
		}
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

func Afficher_Premiers_ID_Relations(allRelations []Globale_Structure_Donnees) {

	// Vérifier que nous avons au moins 2 relations
	if len(allRelations) < ElementsAffichages {
		fmt.Println("Il n'y a pas assez de relations.")
		return
	}

	// Affichage des deux premiers ID des relations
	fmt.Println("Affichage des deux premiers ID de relations :")
	for i := 0; i < 2; i++ {
		relation := allRelations[i]
		fmt.Printf("ID de la Relation: %d\n", relation.ID_RELATIONS)

		// Affichage des lieux et des dates associées pour chaque relation
		for i, lieu := range relation.LOCATIONS_RELATIONS {
			// Affiche le lieu
			fmt.Printf("Lieu: %s\n", lieu)
			// Affiche les dates associées au lieu
			for _, date := range relation.DATESLOCATIONS_RELATIONS[lieu] {
				fmt.Printf("  Date: %s\n", date)
				// fmt.Println("datelocation: ", relation.DATESLOCATIONS_RELATIONS)
			}
			// Séparer les lieux
			if i < len(relation.LOCATIONS_RELATIONS)-1 {
				fmt.Println("------------------------------------------------")
			}
		}
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| AFFICHAGE FINAL |---------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

func Afficher_Premiers_ID_Total(allArtists []Globale_Structure_Donnees, allDates []Globale_Structure_Donnees, allLocations []Globale_Structure_Donnees, allRelations []Globale_Structure_Donnees) {

	// Vérifier qu'il y a au moins deux artistes
	if len(allArtists) < ElementsAffichages {
		fmt.Println("Il n'y a pas assez d'artistes.")
		return
	}

	// Affichage des deux premiers ID des artistes
	fmt.Println("Affichage des deux premiers ID d'artistes :")
	for i := 0; i < 2; i++ {
		artist := allArtists[i]
		date := allDates[i]
		location := allLocations[i]
		relation := allRelations[i]

		fmt.Printf("ID de l'artiste: %d\n", artist.ID_ARTISTE)
		// Affichage des autres informations associées pour chaque ID
		fmt.Println("Nom de l'artiste:", artist.NAME_ARTISTE)
		fmt.Println("Image de l'artiste:", artist.IMAGE)
		fmt.Println("Membres de l'artiste:", artist.MEMBERS_ARTISTE)
		fmt.Println("Date de création de l'artiste:", artist.CREATION_DATE_ARTISTE)
		fmt.Println("Premier album de l'artiste:", artist.FIRST_ALBUM_ARTISTE)
		fmt.Println("Localisations de l'artiste:", location.LOCATIONS_LOCATIONS)
		fmt.Println("Dates des concerts de l'artiste:", date.DATES_DATES)
		fmt.Println("Relations de l'artiste: ")
		// Affichage des lieux et des dates associées pour chaque relation
		for _, lieu := range relation.LOCATIONS_RELATIONS {
			// Affiche le lieu
			fmt.Printf("Lieu: %s\n", lieu)
			// Affiche les dates associées au lieu
			for _, date := range relation.DATESLOCATIONS_RELATIONS[lieu] {
				fmt.Printf("  Date: %s\n", date)
				// fmt.Println("datelocation: ", relation.DATESLOCATIONS_RELATIONS)
			}
			fmt.Println("------------------------------------------------------------------------------------------------------------------")
		}
	}
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
