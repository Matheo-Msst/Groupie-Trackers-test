package main

import (
	"fmt"
	"groupie-tracker/api"
)

func Recup_Donnees() ([]api.Globale_Structure_Donnees, []api.Globale_Structure_Donnees, []api.Globale_Structure_Donnees, []api.Globale_Structure_Donnees) {
	allArtists := api.Api_artists()
	allDates := api.Api_dates()
	allLocations := api.Api_locations()
	allRelations := api.Api_relations()

	return allArtists, allDates, allLocations, allRelations
}

func main() {
	allArtists, allDates, allLocations, allRelations := Recup_Donnees()

	// fmt.Println("//////////////////////////////////////////////////////////////////////////////////////////////////")
	// api.Afficher_Premiers_ID_Artistes(allArtists)
	// fmt.Println("//////////////////////////////////////////////////////////////////////////////////////////////////")
	// api.Afficher_Premiers_ID_Locations(allLocations)
	// fmt.Println("//////////////////////////////////////////////////////////////////////////////////////////////////")
	// api.Afficher_Premiers_ID_Dates(allDates)
	// fmt.Println("//////////////////////////////////////////////////////////////////////////////////////////////////")
	// api.Afficher_Premiers_ID_Relations(allRelations)
	fmt.Println("//////////////////////////////////////////////////////////////////////////////////////////////////")
	api.Afficher_Premiers_ID_Total(allArtists, allDates, allLocations, allRelations)
}
