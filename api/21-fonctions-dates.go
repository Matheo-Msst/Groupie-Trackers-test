package api

import (
	"fmt"
	"strings"
)

// ----------------------------------------------------------------------------------------------------------------------

// Fonction pour enlever les caractères indésirables (_ , -, *)
func Enlever_Caracteres_Dates(TabStr []string) []string {
	var resultat []string

	// Parcours de chaque élément du tableau
	for _, str := range TabStr {
		// Remplacer les caractères "_" et "-" par rien (en les enlevant)
		Str_Traiter := strings.ReplaceAll(str, "_", "")
		Str_Traiter = strings.ReplaceAll(Str_Traiter, "-", "")
		Str_Traiter = strings.ReplaceAll(Str_Traiter, "*", "")
		resultat = append(resultat, Str_Traiter)
	}

	// Retourner les dates nettoyées
	return resultat
}

// ----------------------------------------------------------------------------------------------------------------------

// Fonction pour formater les dates sous la forme "jj/mm/aaaa"
func Triage_Date(dates []string) []string {
	var result []string

	// Pour chaque date dans le tableau
	for _, date := range dates {
		// Vérifier si la date est bien de 8 chiffres (elle sera utilisée après Enlever_Caracteres)
		if len(date) == 8 {
			jour := date[:2]  // Les deux premiers caractères (jour)
			mois := date[2:4] // Les deux suivants (mois)
			annee := date[4:] // Les quatre derniers (année)

			// Formater la date sous forme "jj/mm/aaaa"
			formattedDate := fmt.Sprintf("%s/%s/%s", jour, mois, annee)

			// Ajouter la date formatée au tableau résultat
			result = append(result, formattedDate)
		} else {
			// Si la date n'est pas correcte, ajouter la date brute au tableau
			result = append(result, date)
		}
	}

	// Retourner les dates formatées
	return result
}

// ----------------------------------------------------------------------------------------------------------------------

// Fonction pour récupérer les dates de l'API, les nettoyer et les formater
func Stocker_Dates_Formatees() {
	// Vérifier s'il y a des dates à traiter dans IndexDates
	if len(IndexDates.Dates) == 0 {
		// Pas de dates disponibles dans l'API
		return
	}

	// Initialiser la map pour stocker les dates par ID
	VarFonctions.DATE_FINALE = make(map[int][]string)

	// Pour chaque entrée dans IndexDates, on traite les dates
	for _, dates := range IndexDates.Dates {
		// Appliquer EnleverCaracteres pour enlever les caractères indésirables
		TabDatesModifier := Enlever_Caracteres_Dates(dates.Tab_Dates)

		// Appliquer TriageDate pour formater les dates modifiées
		DATE_FINALE := Triage_Date(TabDatesModifier)

		// Ajouter les dates formatées dans la map sous la clé de l'ID correspondant
		VarFonctions.DATE_FINALE[dates.Id] = DATE_FINALE
	}
}

// ----------------------------------------------------------------------------------------------------------------------

// Fonction pour afficher les dates en sortie d'extraction
func Afficher_Dates() {
	// Vérifier si les données sont présentes
	if len(IndexDates.Dates) == 0 {
		fmt.Println("Aucune donnée de dates disponible.")
		return
	}

	// Afficher les données structurées
	for i, dates := range IndexDates.Dates {
		// Limiter l'affichage en fonction d'un nombre maximal d'éléments
		if i >= ElementAfficher {
			break
		}

		// Affiche l'ID de la date
		fmt.Println("ID:", dates.Id)

		// Afficher chaque date modifiée
		fmt.Println("Dates :")
		for j, lieu := range dates.Tab_Dates {
			fmt.Println(j, ".", lieu)
		}
	}
}

// ----------------------------------------------------------------------------------------------------------------------

// Fonction pour afficher les dates formatées après modification en utilisant VarFonctions
func Afficher_Date_Formatees() {
	// Vérifier s'il y a des dates formatées dans VarFonctions
	if len(VarFonctions.DATE_FINALE) == 0 {
		fmt.Println("Aucune date formatée disponible.")
		return
	}

	// Afficher les données structurées
	for i, dates := range IndexDates.Dates {
		// Limiter l'affichage en fonction d'un nombre maximal d'éléments
		if i >= ElementAfficher {
			break
		}

		// Afficher l'ID de la date
		fmt.Println("\n", "ID:", dates.Id)

		// Afficher chaque date formatée (stockée dans VarFonctions.DATE_FINALE pour cet ID)
		fmt.Println("Dates :")
		for j, date := range VarFonctions.DATE_FINALE[dates.Id] {
			fmt.Println(j, ".", date)
		}
	}
}

// ----------------------------------------------------------------------------------------------------------------------
