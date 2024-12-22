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
	api.Api_Locations()
	api.Afficher_Locations()

	// ------------------------------------------------------------------------------------------------------------//
	fmt.Println("\n //////////////////////////// API DATES //////////////////////////// \n")
	api.Api_dates()

	// Traitement des dates pour enlever les caractères et les formater
	api.Stocker_Dates_Formatees()
	// // Afficher les dates
	// api.Afficher_Dates()

	// Afficher les dates formatées
	api.Afficher_Date_Formatees()

	/* fmt.Println(api.VarFonctions.DATE_FINALE[1][0]) */ //DATE_FINALE[INDEX_DATES][DATE_PRECISE(EXEMPLE ICI LA PREMIER A L'INDEX 0)]

	// ------------------------------------------------------------------------------------------------------------//
	fmt.Println("\n //////////////////////////// API RELATIONS //////////////////////////// \n")
	api.Api_Relations()
	api.Afficher_Relations()

	// ------------------------------------------------------------------------------------------------------------//
}
