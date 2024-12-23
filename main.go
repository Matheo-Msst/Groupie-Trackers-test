package main

import (
	"GroupieTracker/api"
	"fmt"
)

func main() {
	// ------------------------------------------------------------------------------------------------------------//
	fmt.Println("\n //////////////////////////// API ARTISTES //////////////////////////// \n")
	api.Api_Artists()
	api.Afficher_Artistes()

	// ------------------------------------------------------------------------------------------------------------//
	fmt.Println("\n //////////////////////////// API LIEUX //////////////////////////// \n")
	api.Api_Locations() // Récupérer les lieux depuis l'API

	// Traitement et stockage des lieux formatés
	api.Stocker_Lieux_Formates() // Traitement des lieux et stockage dans VarFonctions.LOCATIONS_FINALES

	// Afficher les lieux formatés avec majuscule
	api.Afficher_Lieux_avec_maj() // Affichage des lieux avec la mise en majuscule
	// fmt.Println(api.VarFonctions.LOCATIONS_FINALES)

	// ------------------------------------------------------------------------------------------------------------//
	fmt.Println("\n //////////////////////////// API DATES //////////////////////////// \n")
	api.Api_dates()

	// Traitement des dates pour enlever les caractères et les formater
	api.Stocker_Dates_Formatees()

	api.Afficher_Date_Formatees()
	// fmt.Println(api.VarFonctions.DATE_FINALE)
	// ------------------------------------------------------------------------------------------------------------//
	fmt.Println("\n //////////////////////////// API RELATIONS //////////////////////////// \n")
	api.Api_Relations()
	api.Afficher_Relations()

	// ------------------------------------------------------------------------------------------------------------//
}
